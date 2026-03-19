"""CLI view helpers."""


class CliView:
    def show_menu(self) -> None:
        print("\nВыберите операцию:")
        print("1 - Разбить яйца")
        print("2 - Включить электроплиту")
        print("3 - Добавить масло")
        print("4 - Обжарить яйца")
        print("5 - Добавить приправы")
        print("6 - Перемешать и подать блюдо")
        print("0 - Выход")

    def prompt(self, message: str) -> str:
        return input(message)

    def show_message(self, message: str) -> None:
        print(message)

    def show_error(self, message: str) -> None:
        print(f"Ошибка: {message}")

    def show_state(self, state_text: str) -> None:
        print(f"Состояние: {state_text}")
