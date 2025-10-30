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
 * @author KAnggara75 on Fri 05/09/25 23.25
 * @project PakaiWA pakaiwa_relay
 * https://github.com/PakaiWA/PakaiWA/tree/main/cmd/pakaiwa-relay
 */

package main

import (
	"os"

	"github.com/KAnggara75/scc2go"
)

func init() {
	scc2go.GetEnv(os.Getenv("SCC_URL"), os.Getenv("AUTH"))
}

func main() {}
