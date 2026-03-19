"""PyQt GUI controller."""

from model.cooking_model import CookingModel
from model.exceptions import CookingError
from view.gui_view import QApplication, GuiView


class GuiController:
    def __init__(self, model: CookingModel | None = None, view: GuiView | None = None) -> None:
        self.model = model or CookingModel()
        self.view = view or GuiView()
        self._connect_signals()
        self.update_state()

    def _connect_signals(self) -> None:
        self.view.break_button.clicked.connect(self.break_eggs)
        self.view.stove_button.clicked.connect(self.turn_on_stove)
        self.view.oil_button.clicked.connect(self.add_oil)
        self.view.fry_button.clicked.connect(self.fry_eggs)
        self.view.season_button.clicked.connect(self.add_seasoning)
        self.view.serve_button.clicked.connect(self.stir_and_serve)

    def show(self) -> None:
        self.view.show()

    def break_eggs(self) -> None:
        self._run_action(lambda: self.model.break_eggs(self.view.get_break_amount()))

    def turn_on_stove(self) -> None:
        def action() -> str:
            power, cook_time = self.view.get_stove_settings()
            return f"{self.model.turn_on_stove(power, cook_time)} {self.model.heat_pan()}"

        self._run_action(action)

    def add_oil(self) -> None:
        self._run_action(lambda: self.model.add_oil(self.view.get_oil_type()))

    def fry_eggs(self) -> None:
        self._run_action(lambda: self.model.fry_eggs(self.view.get_fry_amount()))

    def add_seasoning(self) -> None:
        self._run_action(lambda: self.model.add_seasoning(self.view.get_selected_seasonings()))

    def stir_and_serve(self) -> None:
        self._run_action(self.model.stir_and_serve)

    def _run_action(self, action) -> None:
        try:
            result = action()
            self.view.set_status(result)
        except (ValueError, CookingError) as exc:
            self.view.show_error(str(exc))
            self.view.set_status(f"Ошибка: {exc}")
        self.update_state()

    def update_state(self) -> None:
        state = self.model.get_state()
        seasonings = state["seasonings"] or "нет"
        oil_type = state["oil_type"] or "не выбрано"
        state_text = (
            f"Разбито яиц: {state['broken_count']}. "
            f"Обжарено: {state['fried_count']}. "
            f"Сырых осталось: {state['raw_count']}. "
            f"Плита: {'вкл' if state['stove_on'] else 'выкл'} "
            f"(мощность {state['stove_power']}, время {state['cook_time_min']} мин). "
            f"Сковорода: {'разогрета' if state['pan_heated'] else 'холодная'}. "
            f"Масло: {'добавлено' if state['oil_added'] else 'нет'} ({oil_type}). "
            f"Приправы: {seasonings}."
        )
        self.view.set_fry_maximum(int(state["raw_count"]))
        self.view.set_state(state_text)


def run_gui() -> None:
    app = QApplication.instance() or QApplication([])
    controller = GuiController()
    controller.show()
    app.exec()
