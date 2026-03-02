from cooking.exceptions import CookingError
from cooking.process import CookingProcess


def print_menu() -> None:
    print("\nВыберите операцию:")
    print("1 - Разбить яйца")
    print("2 - Включить электроплиту")
    print("3 - Добавить масло")
    print("4 - Обжарить яйца")
    print("5 - Добавить приправы")
    print("6 - Перемешать и подать блюдо")
    print("0 - Выход")


def run_cli() -> None:
    process = CookingProcess()

    while True:
        print_menu()
        choice = input("Введите номер операции: ").strip()

        if choice == "0":
            print("Выход из программы.")
            return

        try:
            if choice == "1":
                amount = int(input("Сколько яиц разбить: ").strip())
                print(process.break_eggs(amount))
                continue

            if choice == "2":
                power = int(input("Выберите мощность конфорки (1-6): ").strip())
                cook_time = int(input("Введите время готовки (мин): ").strip())
                print(process.turn_on_stove(power, cook_time))
                print(process.heat_pan())
                continue

            if choice == "3":
                print("Доступные масла: сливочное, подсолнечное")
                oil_type = input("Какое масло добавить: ").strip()
                print(process.add_oil(oil_type))
                continue

            if choice == "4":
                amount = int(
                    input(
                        f"Сколько яиц обжарить (доступно сырых: {process.eggs.get_raw_count()}): "
                    ).strip()
                )
                print(process.fry_eggs(amount))
                continue

            if choice == "5":
                print("Доступные приправы: соль, перец, паприка, универсальная")
                raw_items = input(
                    "Введите приправы через запятую (например: соль, перец): "
                ).strip()
                items = [item.strip() for item in raw_items.split(",")]
                print(process.add_seasoning(items))
                continue

            if choice == "6":
                print(process.stir_and_serve())
                continue

            print("Ошибка: неизвестный пункт меню.")
        except ValueError:
            print("Ошибка: введите целое число там, где оно требуется.")
        except CookingError as exc:
            print(f"Ошибка: {exc}")
