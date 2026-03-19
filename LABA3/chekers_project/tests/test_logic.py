import unittest
from game_objects import Board, Piece
from game_logic import is_valid_move, check_win_condition

class TestCheckersLogic(unittest.TestCase):
    """
    Набор тестов для проверки ядра игры (Core Logic).
    Тесты запускаются без графического интерфейса Pygame.
    """

    def setUp(self):
        """
        Этот метод автоматически запускается ПЕРЕД каждым тестом.
        Мы создаем абсолютно новую, чистую доску с начальной расстановкой,
        чтобы тесты не влияли друг на друга.
        """
        self.board = Board()

    def test_board_initialization(self):
        """Тест 1: Проверяем правильность создания доски и расстановки."""
        # Проверяем размер доски (8 рядов, 8 колонок)
        self.assertEqual(len(self.board.grid), 8)
        self.assertEqual(len(self.board.grid[0]), 8)
        
        # Проверяем, что на старте в клетке (5, 0) стоит белая шашка
        piece_white = self.board.grid[5][0]
        self.assertIsNotNone(piece_white)
        self.assertEqual(piece_white.color, "white")
        
        # Проверяем, что клетка (4, 0) пустая (там нет шашек на старте)
        self.assertIsNone(self.board.grid[4][0])

    def test_simple_move_validation(self):
        """Тест 2: Проверяем базовые правила ходов (шаг на 1 клетку)."""
        # Белая шашка с (5, 0) хочет пойти по диагонали на (4, 1) - это ПРАВИЛЬНЫЙ ход
        self.assertTrue(is_valid_move(self.board, 5, 0, 4, 1))
        
        # Белая шашка хочет пойти назад на (6, 1) - это НЕПРАВИЛЬНЫЙ ход (она не дамка)
        self.assertFalse(is_valid_move(self.board, 5, 0, 6, 1))
        
        # Шашка хочет пойти прямо вверх по вертикали на (4, 0) - НЕПРАВИЛЬНЫЙ ход (только диагональ)
        self.assertFalse(is_valid_move(self.board, 5, 0, 4, 0))

    def test_jump_logic(self):
        """Тест 3: Проверяем механику рубки (прыжок через шашку)."""
        # Искусственно очищаем доску для нашего сценария
        self.board.grid = [[None for _ in range(8)] for _ in range(8)]
        
        # Ставим нашу белую шашку на (4, 4)
        self.board.grid[4][4] = Piece(4, 4, "white")
        # Ставим вражескую черную шашку на (3, 5)
        self.board.grid[3][5] = Piece(3, 5, "black")
        
        # Проверяем: прыжок через врага на пустую клетку (2, 6) должен быть РАЗРЕШЕН
        self.assertTrue(is_valid_move(self.board, 4, 4, 2, 6))
        
        # Ставим СОЮЗНУЮ белую шашку на (3, 3)
        self.board.grid[3][3] = Piece(3, 3, "white")
        
        # Проверяем: прыжок через свою же шашку на (2, 2) должен быть ЗАПРЕЩЕН
        self.assertFalse(is_valid_move(self.board, 4, 4, 2, 2))

    def test_king_promotion(self):
        """Тест 4: Проверяем механику превращения в дамку."""
        piece = Piece(1, 1, "white")
        
        # При создании шашка НЕ является дамкой
        self.assertFalse(piece.is_king)
        
        # Применяем метод
        piece.make_king()
        
        # Проверяем, что статус изменился
        self.assertTrue(piece.is_king)

    def test_win_condition(self):
        """Тест 5: Проверяем детектор победы."""
        # Искусственно очищаем доску
        self.board.grid = [[None for _ in range(8)] for _ in range(8)]
        
        # Оставляем на доске только одну белую шашку
        self.board.grid[0][0] = Piece(0, 0, "white")
        
        # Вызываем функцию проверки победы. 
        # Так как черных шашек 0, победителями должны быть признаны "white"
        winner = check_win_condition(self.board)
        self.assertEqual(winner, "white")

if __name__ == '__main__':
    unittest.main()