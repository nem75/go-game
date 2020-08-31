package main

import (
        "log"

        "github.com/hajimehoshi/ebiten"
        "github.com/hajimehoshi/ebiten/ebitenutil"
)

func update(screen *ebiten.Image) error {
        if ebiten.IsDrawingSkipped() {
                return nil
        }
        ebitenutil.DebugPrint(screen, "Hello world!")
        return nil
}

func main() {
        if err := ebiten.Run(update, 320, 240, 2, "Ebiten Hello World"); err != nil {
                log.Fatal(err)
        }
}