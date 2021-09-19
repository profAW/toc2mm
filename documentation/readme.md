# toc 2 mm

Tool zur Überführung eines Latex toc-files in ein Plantuml Mindmap

## Architekturgedanken

- einfache Architektur mit prozeduraler Programmierung
- KISS-Prinzip

## Use-Case

- Bus-Systeme: in pptx "stand der lv" zeigen wo wir sind und in Einführungs-LV
- Check-Liste für neue ToDos z.B. in SPS-Pro

## ToDos (simple)

- Error-Handling für z.B. File-Zugriffe
- Unit-Tests anlegen
- direkte Überführung in ein Bild mit hilfe von Plantuml jar-file

## Refactoing

- Paket-Struktur
- Nutzung von Structs wie "Content"
- ggf. mehr OOP Nutzen
- Tests hinzufügen
- Style MM an den Anfang des MM Stellen (Farbe je Ebene), ggf. konfigurierbar über einlesbaren File