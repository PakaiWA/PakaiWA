/*
 * Copyright (c) 2025 KAnggara75
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara75 on Fri 08/08/25 08.29
 * @project PakaiWA pakaiwa
 * https://github.com/KAnggara75/IDXStock/tree/main/cmd/pakaiwa
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("API server berjalan di port 8080...")
	http.ListenAndServe(":8080", nil)
}
