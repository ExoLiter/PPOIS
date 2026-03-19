def is_valid_move(board, start_row, start_col, end_row, end_col):
    """Главная функция проверки валидности хода."""
    piece = board.grid[start_row][start_col]
    if board.grid[end_row][end_col] is not None:
        return False
        
    row_diff = end_row - start_row
    col_diff = end_col - start_col
    
    # Обычный шаг на 1 клетку по диагонали
    if abs(row_diff) == 1 and abs(col_diff) == 1:
        return check_direction(piece, row_diff)
        
    # Прыжок (рубка) на 2 клетки по диагонали
    if abs(row_diff) == 2 and abs(col_diff) == 2:
        return check_jump_logic(board, piece, start_row, start_col, row_diff, col_diff)
        
    return False

def check_direction(piece, row_diff):
    """
    Проверяет правильность направления. Обычные шашки могут двигаться только вперед, дамки - в любом направлении.
    """
    if piece.is_king:
        return True
    if piece.color == "white":
        return row_diff < 0  # Белые всегда идут вверх (индекс строки уменьшается)
    return row_diff > 0      # Черные всегда идут вниз (индекс строки увеличивается)

def check_jump_logic(board, piece, start_row, start_col, row_diff, col_diff):
    """Проверяет наличие врага между начальной и конечной точкой при рубке."""
    if not check_direction(piece, row_diff) and not piece.is_king:
        return False
        
    mid_row = start_row + (row_diff // 2)
    mid_col = start_col + (col_diff // 2)
    mid_piece = board.grid[mid_row][mid_col]
    
    if mid_piece is None:
        return False
        
    # Рубить можно только шашку противоположного цвета
    return mid_piece.color != piece.color

def check_win_condition(board):
    """Проверяет победителя (если не осталось шашек)."""
    white_count = sum(1 for row in board.grid for p in row if p and p.color == "white")
    black_count = sum(1 for row in board.grid for p in row if p and p.color == "black")
    if white_count == 0: return "black"
    if black_count == 0: return "white"
    return None