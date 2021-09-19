# Einführung und Ziele

## Aufgabenstellung
* Grafische Darstellung des TOC eines Latex-Dokumentes als MindMap.
* Umwandlung des toc-files in ein plantuml-mindmap

## Qualitätsziele
* fehlerfreie Umwandlung des TOC in MM (kein Verschieben von Sectionen)
* toc file innerhalb von 5s in MM

## Stakeholder
* Prof. A. Wenzel für seine eigene Orga
* Studierende zum besseren Verständnis der Inhalte und Zusammenhänge

# Randbedingungen 
* Muss auf Windows und Mac laufen
* ggf. Publizierung auf GitHub, wenn alles "chic" ist

# Kontextabgrenzung 

## Fachlicher Kontext
```plantuml
@startuml fachlicherKontext

     !include wzlLib.plantuml

    frame "<&script> Umwandlungsablauf" as Umwandlung{
       
    }
    actor "Anwender" as user
    file "<&image> TOC-MindMap" as mmpng
    collections "<&folder> Latex Projekt" as LatexProjekt


    user --> Umwandlung: steuert
    LatexProjekt <. Umwandlung: nutzt
    Umwandlung --> mmpng: erzeugt

@enduml
```


## Technischer Kontext 
```plantuml
@startuml technischerKontext

     !include wzlLib.plantuml

    frame "<&script> Umwandlungsablauf" as Umwandlung{
        component toc2mm
        file "<&file> TOC-\nMindMap\n.plantuml" as mm
        component "plantuml.jar" as plantuml
        file "<&file> Konfiguration" as config
    }
    actor "Anwender" as user
    file "<&image> TOC-MindMap" as mmpng
    file "<&image> *.toc" as toc
    collections "Latex Projekt" as LatexProjekt


    user --> toc2mm: startet
    config <. toc2mm: nutzt
    toc2mm -.> toc: nutzt
    LatexProjekt .> toc: enthält
    toc2mm -> mm: erzeugt
    mm <. plantuml: nutzt
    user --> plantuml: nutzt
    plantuml --> mmpng: erzeugt

@enduml
```
# Lösungsstrategie
* Umwandlung des TOC (section, subsection, ...) in ein plantuml-mind-map (file.plantuml), dass dann durch ein externes Tool erzeugt werden kann
* Formatierungsvorgagen ggf. über Konfigurationsdatei
* Eigentliche Umwandlung  in ein Bild mittels plantuml.jar und script file

# Bausteinsicht 

## Whitebox Gesamtsystem 

## Ebene 1

### Whitebox *\<Baustein 1\>* 

*\<Whitebox-Template\>*


# Laufzeitsicht 

## *\<Bezeichnung Laufzeitszenario 1\>*

# Verteilungssicht 

## Infrastruktur Ebene 1

# Querschnittliche Konzepte 

## *\<Konzept 1\>* 

# Entwurfsentscheidungen
            * Stategy-Pattern für umsetzung der Anpassung
            * Versuch einer Onion-Struktur

# Qualitätsanforderungen 

## Qualitätsbaum 

## Qualitätsszenarien 


# Risiken und technische Schulden 

# Glossar