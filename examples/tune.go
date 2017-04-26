package main

import (
	"bufio"
	//"fmt"
	"os"
	"time"

	"github.com/ziutek/dvb"
	"github.com/ziutek/dvb/linuxdvb/demux"
	"github.com/ziutek/dvb/linuxdvb/frontend"
	"github.com/ziutek/dvb/ts"
	"github.com/ziutek/dvb/ts/psi"
)

func fail(err error) {
	if err != nil {
		if _, ok := err.(dvb.TemporaryError); !ok {
			panic(err.Error())
		}
	}
}

func main() {
	fe, err := frontend.Open("/dev/dvb/adapter0/frontend0")
	fail(err)

	fail(fe.SetDeliverySystem(dvb.SysATSC))
	fail(fe.SetModulation(dvb.VSB8))
	fail(fe.SetFrequency(uint32(647028615)))
	fail(fe.SetInversion(dvb.InversionAuto))
	fail(fe.Tune())

	deadline := time.Now().Add(5 * time.Second)

	fe3 := frontend.API3{fe}
	var ev frontend.Event
	for ev.Status()&frontend.HasLock == 0 {
		timedout, err := fe3.WaitEvent(&ev, deadline)
		fail(err)
		if timedout {
			println("tuning timeout")
			break
		}
		println(ev.Status())
	}
	println("Done")

	filterParam := demux.StreamFilterParam{
		Pid:  8192,
		In:   demux.InFrontend,
		Out:  demux.OutTSTap,
		Type: demux.Other,
	}

	filter, err := demux.Device("/dev/dvb/adapter0/demux0").NewStreamFilter(&filterParam)
	fail(err)
	fail(filter.Start())

	file, err := os.Open("/dev/dvb/adapter0/dvr0")
	fail(err)
	f := bufio.NewReader(file)

	streamReader := ts.NewPktStreamReader(f)
	decoder := psi.NewSectionDecoder(streamReader, true)

	var pat psi.PAT

	for {
		fail(pat.Update(decoder, true))
		pl := pat.ProgramList()
		for !pl.IsEmpty() {
			pid, pmtid, _ := pl.Pop()
			println("Program", pid, pmtid, pl.IsEmpty())
		}
		/*sl := sdt.ServiceInfo()
		for !sl.IsEmpty() {
			var si psi.ServiceInfo
			si, sl = sl.Pop()
			if si == nil {
				println("Error: damaged service info list")
				break
			}

			dl := si.Descriptors()
			for len(dl) > 0 {
				var d psi.Descriptor
				d, dl = dl.Pop()
				if d == nil {
					println("Error: damaged descriptor list")
					break
				}

				if d.Tag() == psi.ServiceTag {
					sd, ok := psi.ParseServiceDescriptor(d)
					if !ok {
						println("Error: bad service descriptor")
						break
					}
					typ := sd.Type
					name := psi.DecodeText(sd.ServiceName)
					provider := psi.DecodeText(sd.ProviderName)
					fmt.Printf("%v %s %s\n", typ, name, provider)
					break
				}
			}
		}*/
	}
	println("done")
}
