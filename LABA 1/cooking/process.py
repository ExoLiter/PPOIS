from cooking.entities import Eggs, Oil, Pan, Seasoning, Spatula, Stove
from cooking.exceptions import InvalidOperationError


class CookingProcess:
    def __init__(
        self,
        eggs: Eggs | None = None,
        pan: Pan | None = None,
        oil: Oil | None = None,
        seasoning: Seasoning | None = None,
        stove: Stove | None = None,
        spatula: Spatula | None = None,
    ) -> None:
        self.eggs = eggs or Eggs()
        self.pan = pan or Pan()
        self.oil = oil or Oil()
        self.seasoning = seasoning or Seasoning()
        self.stove = stove or Stove()
        self.spatula = spatula or Spatula()
        self.frying_result: str | None = None
        self.recommended_range: tuple[int, int] | None = None

    def break_eggs(self, amount: int) -> str:
        return self.eggs.break_eggs(amount)

    def turn_on_stove(self, power: int, cook_time_min: int) -> str:
        return self.stove.turn_on(power, cook_time_min)

    def heat_pan(self) -> str:
        return self.pan.heat(self.stove)

    def add_oil(self, oil_type: str) -> str:
        return self.oil.add_to_pan(self.pan, oil_type)

    def fry_eggs(self, amount: int) -> str:
        if not self.pan.heated:
            raise InvalidOperationError("Сковорода не разогрета.")
        if not self.oil.added:
            raise InvalidOperationError("Сначала добавьте масло в сковороду.")
        self.eggs.fry(amount)
        self.frying_result, self.recommended_range = self.stove.evaluate_cooking()
        return (
            f"Обжарено яиц: {amount}. "
            f"Всего обжарено: {self.eggs.fried_count}, сырых осталось: {self.eggs.get_raw_count()}. "
            f"Жарка на мощности {self.stove.power} за {self.stove.cook_time_min} мин."
        )

    def add_seasoning(self, items: list[str]) -> str:
        return self.seasoning.add(self.eggs, items)

    def stir_and_serve(self) -> str:
        if not self.seasoning.added_items:
            raise InvalidOperationError("Перед подачей добавьте приправы.")
        if self.frying_result is None or self.recommended_range is None:
            raise InvalidOperationError("Сначала обжарьте яйца.")

        stir_msg = self.spatula.stir(self.eggs)
        serve_msg = self.spatula.serve()
        quality_note = self._build_quality_note()
        seasonings = ", ".join(self.seasoning.added_items)
        return (
            f"{stir_msg} {serve_msg} "
            f"Приправы: {seasonings}. "
            f"Обжарено яиц: {self.eggs.fried_count}, сырых осталось: {self.eggs.get_raw_count()}. "
            f"{quality_note}"
        )

    def _build_quality_note(self) -> str:
        if self.frying_result is None or self.recommended_range is None:
            raise InvalidOperationError("Результат жарки не определен.")
        if self.frying_result == "raw":
            return (
                "Яйца получились сырые. "
                f"Нужно было жарить {self.recommended_range[0]}-{self.recommended_range[1]} мин."
            )
        if self.frying_result == "burned":
            return (
                "Яйца сгорели. "
                f"Нужно было жарить {self.recommended_range[0]}-{self.recommended_range[1]} мин."
            )
        return "Яйца приготовлены правильно."
