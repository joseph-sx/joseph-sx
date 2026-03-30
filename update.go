package main

import (
	"log"
	"math"
	"os"
	"strings"
	"text/template"
	"time"
)

type ProfileData struct {
	Year            int
	ProgressBar     string
	ProgressPercent float64
	LastUpdated     string
}

func main() {
	t := time.Now()
	thisYear := t.Year()
	startTime := time.Date(thisYear, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
	endTime := time.Date(thisYear+1, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano()
	now := t.UnixNano()

	progress := float64(now-startTime) / float64(endTime-startTime)
	capacity := 30.0
	passed := int(math.Floor(progress * capacity))
	bar := strings.Repeat("■", passed) + strings.Repeat("□", int(capacity)-passed)
	bt := "`"

	const readmeTmpl = `
# 🐙 Joseph García (Joe)
### Systems Engineer | DevOps & Automation Specialist | Go Developer

Senior Systems Engineer con +10 años de trayectoria. Especializado en el diseño de infraestructuras escalables, automatización de procesos y microservicios robustos. Actualmente protegiendo activos digitales en **HSBC México**.

---

### 🏗️ The Kraken Hub (Personal Infrastructure)
Mantengo mi propio ecosistema de **Self-hosting** bajo la filosofía de soberanía digital:
* **Core:** Linux (Debian/Ubuntu/AlpineLinux), Docker, HAProxy, WireGuard.
* **Management:** Portainer, CI/CD con GitHub Actions.
* **Stack:** Microservicios en **Go** y automatización de Telegram Bots.

### 🛠️ Tecnologías & Core Stack
* **Lenguajes:** {{.BT}}Go (Golang){{.BT}}, {{.BT}}Dart (Flutter){{.BT}}, {{.BT}}SQL{{.BT}}, {{.BT}}Bash{{.BT}}, {{.BT}}JS/TS{{.BT}}, {{.BT}}Python{{.BT}} {{.BT}}PHP{{.BT}}.
* **Infraestructura:** {{.BT}}Docker{{.BT}}, {{.BT}}Self-hosting{{.BT}}, {{.BT}}API Design{{.BT}}, {{.BT}}Data Modeling{{.BT}}.
* **Editor:** {{.BT}}vi{{.BT}} (porque la eficiencia nace en la terminal).

### 🔭 Proyectos Destacados
* 🛰️ **StreamBot:** Microservicio en **Go** + Streamlink para gestión de contenido multimedia.
* 📦 **SaaS Multitenant:** Arquitectura distribuida con panel OTP para gestión de comunidades.
* 🔌 **WP Eco-System:** Plugins de WooCommerce con despliegues continuos y self-hosted updates.

---

### 📈 Estadísticas & Actividad
![Top Langs](https://gh-stats-cards.vercel.app/api/top-langs/?username=joseph-sx&layout=donut&langs_count=10&theme=dark&hide=html,css,scss)

---

### 💬 Contacto & Consultoría
* **Especialidades:** Seguridad, modelado de datos y arquitectura de sistemas en Go.
* **Freelance:** [The Kraken Hub Ops en Fiverr](https://www.fiverr.com/joseph_sx) - DevOps & Systems Engineering.
* **Blog:** *Ideas en cola:* Despliegues Zero-Downtime y el arte del Self-hosting.

---

> "Build for scale, automate for peace of mind."

⏳ **Year Progress {{.Year}}**
[{{.ProgressBar}}]  {{printf "%.2f" .ProgressPercent}} %%
*Last updated: {{.LastUpdated}}*

![](https://komarev.com/ghpvc/?username=joseph-sx&label=PROFILE+VIEWS)
`
	data := struct {
		ProfileData
		BT string
	}{
		ProfileData: ProfileData{
			Year:            thisYear,
			ProgressBar:     bar,
			ProgressPercent: progress * 100,
			LastUpdated:     t.Format("2006-01-02"),
		},
		BT: bt,
	}

	// 5. Escritura atómica del archivo
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatalf("Error creando archivo: %v", err)
	}
	defer f.Close()

	tmpl, err := template.New("readme").Parse(readmeTmpl)
	if err != nil {
		log.Fatalf("Error en template: %v", err)
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		log.Fatalf("Error ejecutando render: %v", err)
	}
}
