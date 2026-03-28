package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	thisYear := time.Now().Year()

	// Lógica de progreso del año
	startTimeOfThisYear := time.Date(thisYear, time.January, 1, 0, 0, 0, 0, time.UTC).UnixNano()
	endTimeOfThisYear := time.Date(thisYear, time.December, 31, 23, 59, 59, 0, time.UTC).UnixNano()

	progressOfThisYear := float64(time.Now().UnixNano()-startTimeOfThisYear) /
		float64(endTimeOfThisYear-startTimeOfThisYear)

	progressBarCapacity := 30.0
	passedProgressBarIndex := int(math.Floor(progressOfThisYear * progressBarCapacity))

	passedProgressBar := ""
	for i := 0; i < passedProgressBarIndex; i++ {
		passedProgressBar += "■"
	}

	leftProgressBar := ""
	for i := 0; i < int(progressBarCapacity)-int(passedProgressBarIndex); i++ {
		leftProgressBar += "□"
	}

	file, err := os.Create("README.md")
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()

	// Definición del nuevo contenido profesional
	README := "# 🐙 Joseph García (Joe)\n" +
		"### Systems Engineer | DevOps & Automation Specialist | Go Developer\n\n" +
		"Senior Systems Engineer con +10 años de trayectoria. Especializado en el diseño de infraestructuras escalables, automatización de procesos y microservicios robustos. Actualmente protegiendo activos digitales en **HSBC México**.\n\n" +
		"---\n\n" +
		"### 🏗️ The Kraken Hub (Personal Infrastructure)\n" +
		"Mantengo mi propio ecosistema de **Self-hosting** bajo la filosofía de soberanía digital:\n" +
		"* **Core:** Linux (Debian/Ubuntu/AlpineLinux), Docker, HAProxy, WireGuard.\n" +
		"* **Management:** Portainer, CI/CD con GitHub Actions.\n" +
		"* **Stack:** Microservicios en **Go** y automatización de Telegram Bots.\n\n" +
		"### 🛠️ Tecnologías & Core Stack\n" +
		"* **Lenguajes:** `Go (Golang)`, `Dart (Flutter)`, `SQL`, `Bash`, `JS/TS`, `Python` `PHP`.\n" +
		"* **Infraestructura:** `Docker`, `Self-hosting`, `API Design`, `Data Modeling`.\n" +
		"* **Editor:** `vi` (porque la eficiencia nace en la terminal).\n\n" +
		"### 🔭 Proyectos Destacados\n" +
		"* 🛰️ **StreamBot:** Microservicio en **Go** + Streamlink para gestión de contenido multimedia.\n" +
		"* 📦 **SaaS Multitenant:** Arquitectura distribuida con panel OTP para gestión de comunidades.\n" +
		"* 🔌 **WP Eco-System:** Plugins de WooCommerce con despliegues continuos y self-hosted updates.\n\n" +
		"---\n\n" +
		"### 📈 Estadísticas & Actividad\n" +
		"![Top Langs](https://gh-stats-cards.vercel.app/api/top-langs/?username=joseph-sx&layout=donut&langs_count=10&theme=dark&hide=html,css,scss)\n\n" +
		"---\n\n" +
		"### 💬 Contacto & Consultoría\n" +
		"* **Especialidades:** Seguridad, modelado de datos y arquitectura de sistemas en Go.\n" +
		"* **Freelance:** [The Kraken Hub Ops en Fiverr](https://www.fiverr.com/joseph_sx) - DevOps & Systems Engineering.\n" +
		"* **Blog:** *Ideas en cola:* Despliegues Zero-Downtime y el arte del Self-hosting.\n\n" +
		"---\n\n" +
		"> \"Build for scale, automate for peace of mind.\"\n\n" +
		"⏳ **Year Progress " + strconv.Itoa(thisYear) + "**\n" +
		"[" + passedProgressBar + leftProgressBar + "]  " + strconv.FormatFloat(progressOfThisYear*100, 'f', 2, 64) + " %\n" +
		"*Last updated: " + currentDate + "*"

	_, err = file.WriteString(README)
	if err != nil {
		log.Fatalf("Failed write to file: %s", err)
	}
}
