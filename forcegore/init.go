package forcegore

import eggosLog "github.com/icexin/eggos/log"

func init() {
	err := eggosLog.SetLevel(eggosLog.LoglvlDebug)
	if err != nil {
		panic(err)
	}
}
