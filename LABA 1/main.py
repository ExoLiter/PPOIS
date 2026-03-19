"""Program entrypoint."""

import sys

from controller.cli_controller import CliController


def main() -> None:
    if "--gui" in sys.argv:
        try:
            from controller.gui_controller import run_gui
        except ImportError as exc:
            raise SystemExit(
                "Для запуска GUI установите PyQt5 или PyQt6."
            ) from exc

        run_gui()
        return
    CliController().run()


if __name__ == "__main__":
    main()
