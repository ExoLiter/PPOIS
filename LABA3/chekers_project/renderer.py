import pygame
from settings import CONFIG

def draw_game_surface(board, selected, anim_data):
    """Отрисовывает логическую доску (без переворота) на отдельную поверхность."""
    size = CONFIG["window"]["width"] # Предполагаем квадратное окно, ширина = высота
    surface = pygame.Surface((size, size)) 
    draw_board_grid(surface) # Рисуем клеточную сетку
    
    if selected: 
        draw_highlight(surface, selected[0], selected[1]) # Подсвечиваем выбранную шашку
        
    draw_pieces(surface, board, anim_data) 
    draw_animation(surface, anim_data) 
    return surface

def draw_rotated_screen(screen, game_surface, angle):
    """Плавно вращает доску и выводит на экран."""
    bg_color = CONFIG["colors"]["bg_menu"] # Фон для заполнения пустых углов при вращении
    screen.fill(bg_color) 
    
    rotated = pygame.transform.rotate(game_surface, angle) # Поворачиваем поверхность с игрой на заданный угол
    rect = rotated.get_rect(center=(400, 400)) # Центрируем повернутую поверхность в окне (400, 400 - центр окна 800x800)
    screen.blit(rotated, rect.topleft) # Рисуем повернутую поверхность на экране, используя координаты верхнего левого угла для позиционирования

def draw_board_grid(surface):
    """Рисует клеточную сетку."""
    cell_size = CONFIG["grid"]["cell_size"] # Размер клетки, например 100 для 8x8 на 800x800
    for r in range(8): 
        for c in range(8):  
            color = CONFIG["colors"]["board_dark"] if (r+c) % 2 != 0 else CONFIG["colors"]["board_light"] # Чередуем цвета для клеток
            rect = pygame.Rect(c * cell_size, r * cell_size, cell_size, cell_size) # Создаем прямоугольник для клетки
            pygame.draw.rect(surface, color, rect) # Рисуем клетку на поверхности

def draw_highlight(surface, row, col):
    """Подсвечивает выбранную шашку (желтый квадрат)."""
    cell_size = CONFIG["grid"]["cell_size"]
    color = CONFIG["colors"]["highlight"]
    rect = pygame.Rect(col * cell_size, row * cell_size, cell_size, cell_size)
    pygame.draw.rect(surface, color, rect, 5)

def draw_pieces(surface, board, active_anim):
    """Отрисовывает статические шашки."""
    for r in range(8):
        for c in range(8):
            piece = board.grid[r][c]
            if piece and not (active_anim["active"] and active_anim["piece"] == piece):
                draw_single_piece(surface, piece, c * 100, r * 100)

def draw_single_piece(surface, piece, x_pos, y_pos):
    """Рисует кружок шашки."""
    center = (int(x_pos + 50), int(y_pos + 50))
    color = CONFIG["colors"]["piece_" + piece.color]
    pygame.draw.circle(surface, color, center, 40)
    
    if piece.is_king:
        crown = CONFIG["colors"]["king_crown"]
        pygame.draw.circle(surface, crown, center, 20)

def draw_animation(surface, anim_data):
    """Рисует шашку в полете во время анимации."""
    if not anim_data["active"]: return
    prog = anim_data["progress"]
    
    start_x, start_y = anim_data["start_col"] * 100, anim_data["start_row"] * 100
    end_x, end_y = anim_data["end_col"] * 100, anim_data["end_row"] * 100
    
    cur_x = start_x + (end_x - start_x) * prog
    cur_y = start_y + (end_y - start_y) * prog
    draw_single_piece(surface, anim_data["piece"], cur_x, cur_y)