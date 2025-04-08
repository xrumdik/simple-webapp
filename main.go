package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, WORLD, 世界, МИР, PEACE, FRIEDEN, PAIX, PACE, PAZ, VREDE, POKÓJ, MÍR, MIER, МІР, MIR, BÉKE, PACE, ΕΙΡΉΝΗ, BARIŞ, RAUHA, FRED, FRED, FRED, RAHU, MIERS, TAIKA, 和平, 平和, 평화, HÒA BÌNH, शांति, سلام, שלום, صلح, امن, สันติภาพ, ЭНХ ТАЙВАН, მშვიდობა, ԽԱՂԱՂՈՒԹՅՈՒՆ, БЕЙБІТШІЛІК, TINCHLIK, ТЫНЫЧЛЫК, ТЫНЫС, АҺАР, ТЫНЧТЫК, PARAHATÇYLYK, СУЛҲ, SÜLH!!!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running simple-webapp app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
