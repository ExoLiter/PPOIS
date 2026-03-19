from settings import CONFIG

class Piece:
    """
    Класс описывает шашку.
    Хранит ее цвет, статус дамки и логические координаты.
    """
    def __init__(self, row, col, color):
        """Инициализация новой шашки."""
        self.row = row # Логические координаты, не зависят от размера клеток
        self.col = col # Логические координаты, не зависят от размера клеток
        self.color = color # "white" или "black"
        self.is_king = False # Статус дамки, по умолчанию False

    def make_king(self):
        """Превращает обычную шашку в дамку."""
        self.is_king = True

class Board:
    """
    Класс доски. Хранит двумерный массив шашек.
    """
    def __init__(self):
        """Инициализирует пустую доску и расставляет шашки."""
        self.grid = []
        self.rows = CONFIG["grid"]["rows"] 
        self.cols = CONFIG["grid"]["cols"]
        self.setup_board()

    def setup_board(self):
        """Расставляет шашки по правилам английских шашек."""
        for row in range(self.rows):
            self.grid.append([])
            for col in range(self.cols):
                piece = self.create_piece_for_cell(row, col)
                self.grid[row].append(piece)

    def create_piece_for_cell(self, row, col):
        """Определяет, какая шашка должна стоять в ячейке при старте."""
        is_dark_cell = (row + col) % 2 != 0 # Темные клетки - те, на которых стоят шашки
        if not is_dark_cell: 
            return None # На светлых клетках нет шашек
            
        if row < 3:
            return Piece(row, col, "black") # Черные стоят вверху (меньшие индексы строк)
        if row > 4:
            return Piece(row, col, "white") # Белые стоят внизу (большие индексы строк)
            
        return None

    def move_piece(self, start_row, start_col, end_row, end_col):
        """Перемещает шашку в массиве."""
        piece = self.grid[start_row][start_col] # Получаем шашку, которую нужно переместить
        self.grid[start_row][start_col] = None # Очищаем старую позицию
        self.grid[end_row][end_col] = piece # Ставим шашку на новую позицию
        
        piece.row = end_row # Обновляем логические координаты шашки
        piece.col = end_col

    def remove_piece(self, row, col):
        """Удаляет шашку с доски (при рубке).""" 
        self.grid[row][col] = None # Просто ставим None, чтобы обозначить пустую клетку