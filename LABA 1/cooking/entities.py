from cooking.exceptions import InvalidOperationError


class Eggs:
    def __init__(
        self,
        broken_count: int = 0,
        fried_count: int = 0,
        stirred: bool = False,
        max_broken_count: int = 10,
    ) -> None:
        self.broken_count = broken_count
        self.fried_count = fried_count
        self.stirred = stirred
        self.max_broken_count = max_broken_count

    def break_eggs(self, amount: int) -> str:
        if self.fried_count > 0:
            raise InvalidOperationError("Нельзя добавлять яйца после обжаривания.")
        if amount <= 0:
            raise InvalidOperationError("Количество яиц должно быть больше нуля.")
        if self.broken_count + amount > self.max_broken_count:
            raise InvalidOperationError("Максимум можно разбить 10 яиц.")
        self.broken_count += amount
        return f"Добавлено {amount} яиц. Всего разбито яиц: {self.broken_count}."

    def fry(self, amount: int) -> None:
        if self.broken_count <= 0:
            raise InvalidOperationError("Сначала нужно разбить яйца.")
        if amount <= 0:
            raise InvalidOperationError("Количество обжариваемых яиц должно быть больше нуля.")
        if amount > self.get_raw_count():
            raise InvalidOperationError("Нельзя обжарить больше яиц, чем доступно сырых.")
        self.fried_count += amount

    def stir(self) -> str:
        if self.fried_count <= 0:
            raise InvalidOperationError("Сначала нужно обжарить яйца.")
        if self.stirred:
            raise InvalidOperationError("Яичница уже перемешана.")
        self.stirred = True
        return "Яичница перемешана лопаткой."

    def get_raw_count(self) -> int:
        return self.broken_count - self.fried_count


class Stove:
    def __init__(
        self,
        is_on: bool = False,
        power: int | None = None,
        cook_time_min: int | None = None,
        ranges: dict[int, tuple[int, int]] | None = None,
    ) -> None:
        self.is_on = is_on
        self.power = power
        self.cook_time_min = cook_time_min
        self.ranges = ranges or {
            1: (15, 20),
            2: (11, 14),
            3: (7, 10),
            4: (5, 6),
            5: (3, 4),
            6: (1, 2),
        }

    def turn_on(self, power: int, cook_time_min: int) -> str:
        if power not in self.ranges:
            raise InvalidOperationError("Мощность должна быть в диапазоне от 1 до 6.")
        if cook_time_min <= 0:
            raise InvalidOperationError("Время готовки должно быть больше нуля.")
        self.is_on = True
        self.power = power
        self.cook_time_min = cook_time_min
        return (
            f"Электроплита включена. Мощность конфорки: {power}, "
            f"время готовки: {cook_time_min} мин."
        )

    def evaluate_cooking(self) -> tuple[str, tuple[int, int]]:
        if not self.is_on or self.power is None or self.cook_time_min is None:
            raise InvalidOperationError("Сначала включите электроплиту и задайте параметры.")
        low, high = self.ranges[self.power]
        if self.cook_time_min < low:
            return "raw", (low, high)
        if self.cook_time_min > high:
            return "burned", (low, high)
        return "good", (low, high)


class Pan:
    def __init__(self, heated: bool = False) -> None:
        self.heated = heated

    def heat(self, stove: Stove) -> str:
        if not stove.is_on:
            raise InvalidOperationError("Плита выключена. Сначала включите электроплиту.")
        if self.heated:
            raise InvalidOperationError("Сковорода уже разогрета.")
        self.heated = True
        return "Сковорода разогрета."


class Oil:
    def __init__(
        self,
        oil_type: str | None = None,
        added: bool = False,
        allowed_types: set[str] | None = None,
    ) -> None:
        self.oil_type = oil_type
        self.added = added
        self.allowed_types = allowed_types or {"сливочное", "подсолнечное"}

    def add_to_pan(self, pan: Pan, oil_type: str) -> str:
        if not pan.heated:
            raise InvalidOperationError("Сначала разогрейте сковороду.")
        normalized = oil_type.strip().lower()
        if normalized not in self.allowed_types:
            raise InvalidOperationError("Доступны только: сливочное или подсолнечное масло.")
        if self.added:
            raise InvalidOperationError("Масло уже добавлено в сковороду.")
        self.oil_type = normalized
        self.added = True
        return f"Добавлено {normalized} масло."


class Seasoning:
    def __init__(
        self,
        allowed: set[str] | None = None,
        added_items: list[str] | None = None,
    ) -> None:
        self.allowed = allowed or {"соль", "перец", "паприка", "универсальная"}
        self.added_items = added_items or []

    def add(self, eggs: Eggs, items: list[str]) -> str:
        if eggs.fried_count <= 0:
            raise InvalidOperationError("Приправы добавляются после обжаривания.")
        if not items:
            raise InvalidOperationError("Нужно выбрать хотя бы одну приправу.")

        normalized_items: list[str] = []
        for item in items:
            normalized = item.strip().lower()
            if not normalized:
                continue
            if normalized not in self.allowed:
                raise InvalidOperationError(
                    "Недопустимая приправа. Можно: соль, перец, паприка, универсальная."
                )
            if normalized not in self.added_items:
                self.added_items.append(normalized)
                normalized_items.append(normalized)

        if not normalized_items:
            raise InvalidOperationError("Новые приправы не выбраны.")
        return f"Добавлены приправы: {', '.join(normalized_items)}."


class Spatula:
    def __init__(self, clean: bool = True, used: bool = False) -> None:
        self.clean = clean
        self.used = used

    def stir(self, eggs: Eggs) -> str:
        self.used = True
        self.clean = False
        return eggs.stir()

    def serve(self) -> str:
        if not self.used:
            raise InvalidOperationError("Сначала перемешайте блюдо лопаткой.")
        return "Блюдо подано."
