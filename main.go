package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world, 世界, мир, peace, frieden, paix, pace, paz, vrede, pokój, mír, mier, мір, mir, béke, pace, ειρήνη, barış, rauha, fred, fred, fred, rahu, miers, taika, 和平, 平和, 평화, hòa bình, शांति, سلام, שלום, صلح, امن, สันติภาพ, энх тайван, მშვიდობა, խաղաղություն, бейбітшілік, tinchlik, тынычлык, тыныс, аhар, тынчтык, parahatçylyk, сулҳ, sülh")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running simple-webapp app. Press Ctrl+C to exit...")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
