import unittest

from cooking.exceptions import InvalidOperationError
from cooking.process import CookingProcess


class TestCookingProcess(unittest.TestCase):
    def test_happy_path(self) -> None:
        process = CookingProcess()
        process.break_eggs(3)
        process.turn_on_stove(4, 6)
        process.heat_pan()
        process.add_oil("сливочное")
        process.fry_eggs(2)
        process.add_seasoning(["соль", "перец"])
        result = process.stir_and_serve()
        self.assertIn("Яйца приготовлены правильно", result)
        self.assertIn("соль", result)
        self.assertIn("Обжарено яиц: 2", result)
        self.assertIn("сырых осталось: 1", result)

    def test_break_eggs_accumulates(self) -> None:
        process = CookingProcess()
        process.break_eggs(3)
        process.break_eggs(2)
        self.assertEqual(process.eggs.broken_count, 5)

    def test_break_eggs_has_max_limit(self) -> None:
        process = CookingProcess()
        process.break_eggs(10)
        with self.assertRaises(InvalidOperationError):
            process.break_eggs(1)

    def test_stove_invalid_power(self) -> None:
        process = CookingProcess()
        with self.assertRaises(InvalidOperationError):
            process.turn_on_stove(7, 5)

    def test_add_oil_before_heating_pan(self) -> None:
        process = CookingProcess()
        with self.assertRaises(InvalidOperationError):
            process.add_oil("подсолнечное")

    def test_serve_reports_burned_eggs_with_recommended_time(self) -> None:
        process = CookingProcess()
        process.break_eggs(2)
        process.turn_on_stove(6, 1)
        process.heat_pan()
        process.add_oil("подсолнечное")
        process.fry_eggs(2)
        process.add_seasoning(["паприка"])
        result = process.stir_and_serve()
        self.assertIn("Яйца приготовлены правильно", result)

        burned = CookingProcess()
        burned.break_eggs(2)
        burned.turn_on_stove(6, 5)
        burned.heat_pan()
        burned.add_oil("подсолнечное")
        burned.fry_eggs(2)
        burned.add_seasoning(["паприка"])
        burned_result = burned.stir_and_serve()
        self.assertIn("Яйца сгорели", burned_result)
        self.assertIn("1-2 мин", burned_result)

    def test_serve_reports_raw_eggs_with_recommended_time(self) -> None:
        process = CookingProcess()
        process.break_eggs(2)
        process.turn_on_stove(4, 2)
        process.heat_pan()
        process.add_oil("подсолнечное")
        process.fry_eggs(2)
        process.add_seasoning(["соль"])
        result = process.stir_and_serve()
        self.assertIn("Яйца получились сырые", result)
        self.assertIn("5-6 мин", result)

    def test_cannot_serve_without_seasoning(self) -> None:
        process = CookingProcess()
        process.break_eggs(2)
        process.turn_on_stove(4, 5)
        process.heat_pan()
        process.add_oil("сливочное")
        process.fry_eggs(2)
        with self.assertRaises(InvalidOperationError):
            process.stir_and_serve()

    def test_cannot_fry_more_than_raw(self) -> None:
        process = CookingProcess()
        process.break_eggs(6)
        process.turn_on_stove(4, 5)
        process.heat_pan()
        process.add_oil("сливочное")
        with self.assertRaises(InvalidOperationError):
            process.fry_eggs(7)


if __name__ == "__main__":
    unittest.main()
