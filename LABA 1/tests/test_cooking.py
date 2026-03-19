import unittest

from controller.cli_controller import CliController
from model.cooking_model import CookingModel
from model.exceptions import InvalidOperationError
from view.cli_view import CliView


class DummyCliView(CliView):
    def __init__(self) -> None:
        self.messages: list[str] = []
        self.errors: list[str] = []
        self.states: list[str] = []

    def prompt(self, message: str) -> str:
        raise NotImplementedError

    def show_message(self, message: str) -> None:
        self.messages.append(message)

    def show_error(self, message: str) -> None:
        self.errors.append(message)

    def show_state(self, state_text: str) -> None:
        self.states.append(state_text)

    def show_menu(self) -> None:
        return


class TestCookingModel(unittest.TestCase):
    def test_happy_path(self) -> None:
        model = CookingModel()
        model.break_eggs(3)
        model.turn_on_stove(4, 6)
        model.heat_pan()
        model.add_oil("сливочное")
        model.fry_eggs(2)
        model.add_seasoning(["соль", "перец"])
        result = model.stir_and_serve()
        self.assertIn("Яйца приготовлены правильно", result)
        self.assertIn("соль", result)
        self.assertIn("Обжарено яиц: 2", result)
        self.assertIn("сырых осталось: 1", result)

    def test_break_eggs_accumulates(self) -> None:
        model = CookingModel()
        model.break_eggs(3)
        model.break_eggs(2)
        self.assertEqual(model.eggs.broken_count, 5)

    def test_break_eggs_has_max_limit(self) -> None:
        model = CookingModel()
        model.break_eggs(10)
        with self.assertRaises(InvalidOperationError):
            model.break_eggs(1)

    def test_stove_invalid_power(self) -> None:
        model = CookingModel()
        with self.assertRaises(InvalidOperationError):
            model.turn_on_stove(7, 5)

    def test_add_oil_before_heating_pan(self) -> None:
        model = CookingModel()
        with self.assertRaises(InvalidOperationError):
            model.add_oil("подсолнечное")

    def test_serve_reports_burned_eggs_with_recommended_time(self) -> None:
        model = CookingModel()
        model.break_eggs(2)
        model.turn_on_stove(6, 5)
        model.heat_pan()
        model.add_oil("подсолнечное")
        model.fry_eggs(2)
        model.add_seasoning(["паприка"])
        result = model.stir_and_serve()
        self.assertIn("Яйца сгорели", result)
        self.assertIn("1-2 мин", result)

    def test_serve_reports_raw_eggs_with_recommended_time(self) -> None:
        model = CookingModel()
        model.break_eggs(2)
        model.turn_on_stove(4, 2)
        model.heat_pan()
        model.add_oil("подсолнечное")
        model.fry_eggs(2)
        model.add_seasoning(["соль"])
        result = model.stir_and_serve()
        self.assertIn("Яйца получились сырые", result)
        self.assertIn("5-6 мин", result)

    def test_cannot_serve_without_seasoning(self) -> None:
        model = CookingModel()
        model.break_eggs(2)
        model.turn_on_stove(4, 5)
        model.heat_pan()
        model.add_oil("сливочное")
        model.fry_eggs(2)
        with self.assertRaises(InvalidOperationError):
            model.stir_and_serve()

    def test_cannot_fry_more_than_raw(self) -> None:
        model = CookingModel()
        model.break_eggs(6)
        model.turn_on_stove(4, 5)
        model.heat_pan()
        model.add_oil("сливочное")
        with self.assertRaises(InvalidOperationError):
            model.fry_eggs(7)

    def test_get_state_returns_current_snapshot(self) -> None:
        model = CookingModel()
        model.break_eggs(4)
        state = model.get_state()
        self.assertEqual(state["broken_count"], 4)
        self.assertEqual(state["raw_count"], 4)
        self.assertFalse(bool(state["stove_on"]))


class TestCliController(unittest.TestCase):
    def test_format_state_uses_model_snapshot(self) -> None:
        model = CookingModel()
        model.break_eggs(3)
        view = DummyCliView()
        controller = CliController(model=model, view=view)
        state_text = controller.format_state()
        self.assertIn("Разбито: 3", state_text)
        self.assertIn("сырых осталось: 3", state_text)


if __name__ == "__main__":
    unittest.main()
