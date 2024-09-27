# GoCalculator

GoCalculator ist eine einfache Konsolenanwendung in Go,
 die grundlegende arithmetische Operationen ermöglicht.
 Dieses Projekt dient als Beispiel, um die Funktionsweise 
von GitLab CI/CD zu demonstrieren.

## Funktionen

- Addition
- Subtraktion
- Multiplikation
- Division
- Prozentrechnung

## tests
7 Unittests

## Voraussetzungen

Die Anwendung wurde mit Go 1.13 entwickelt und getestet.

## Lokale Installation

Um das Projekt lokal zu installieren, 
klone das Repository und führe die Anwendung aus:



## zum aufbauen:

mkdir GoCalculator
cd GoCalculator

go mod init GoCalculator 
go mod tidy 



## zum ausführen

```bash
cd GoCalculator
go run calculator.go
```


## ausführen der tests:

```bash
go test
```