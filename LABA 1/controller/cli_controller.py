"""CLI controller."""

from model.cooking_model import CookingModel
from model.exceptions import CookingError 
from view.cli_view import CliView 


class CliController:
    def __init__(self, model: CookingModel | None = None, view: CliView | None = None) -> None:
        self.model = model or CookingModel() # Инициализация модели, если не передана
        self.view = view or CliView()

    def run(self) -> None:
        while True:
            self.view.show_menu()
            choice = self.view.prompt("Введите номер операции: ").strip()
            if choice == "0":
                self.view.show_message("Выход из программы.")
                return
            self._handle_choice(choice)

    def _handle_choice(self, choice: str) -> None:
        try:
            if choice == "1":
                amount = int(self.view.prompt("Сколько яиц разбить: ").strip())
                result = self.model.break_eggs(amount)
            elif choice == "2":
                power = int(self.view.prompt("Выберите мощность конфорки (1-6): ").strip())
                cook_time = int(self.view.prompt("Введите время готовки (мин): ").strip())
                result = f"{self.model.turn_on_stove(power, cook_time)} {self.model.heat_pan()}"
            elif choice == "3":
                oil_type = self.view.prompt(
                    "Какое масло добавить (сливочное/подсолнечное): "
                ).strip()
                result = self.model.add_oil(oil_type)
            elif choice == "4":
                amount = int(
                    self.view.prompt(
                        f"Сколько яиц обжарить (доступно сырых: {self.model.eggs.get_raw_count()}): "
                    ).strip()
                )
                result = self.model.fry_eggs(amount)
            elif choice == "5":
                raw_items = self.view.prompt(
                    "Введите приправы через запятую (соль, перец, паприка, универсальная): "
                ).strip()
                items = [item.strip() for item in raw_items.split(",")]
                result = self.model.add_seasoning(items)
            elif choice == "6":
                result = self.model.stir_and_serve()
            else:
                self.view.show_error("неизвестный пункт меню")
                return

            self.view.show_message(result)
            self.view.show_state(self.format_state())
        except (ValueError, CookingError) as exc:
            self.view.show_error(str(exc))

    def format_state(self) -> str:
        state = self.model.get_state()
        seasonings = state["seasonings"] or "нет"
        oil_type = state["oil_type"] or "не выбрано"
        return (
            f"Разбито: {state['broken_count']}, обжарено: {state['fried_count']}, "
            f"сырых осталось: {state['raw_count']}. "
            f"Плита: {'вкл' if state['stove_on'] else 'выкл'}, "
            f"мощность: {state['stove_power']}, время: {state['cook_time_min']} мин. "
            f"Сковорода: {'разогрета' if state['pan_heated'] else 'холодная'}. "
            f"Масло: {'добавлено' if state['oil_added'] else 'нет'} ({oil_type}). "
            f"Приправы: {seasonings}."
        )
