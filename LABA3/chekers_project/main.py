import pygame
import sys
from settings import CONFIG, get_text
from audio_manager import initialize_audio, play_sound, adjust_volume, stop_music
from records_manager import load_records, save_winner
from game_objects import Board
from game_logic import is_valid_move, check_win_condition
from renderer import draw_game_surface, draw_rotated_screen
from ui_components import draw_menu_layout, handle_menu_click, draw_text, draw_multiline_text, draw_overlay

def run_menu_state(screen):
    """Главное меню."""
    clock = pygame.time.Clock()
    while True:
        btns = [
            {"key": "play", "y": 300, "action": "PLAY"},
            {"key": "records", "y": 400, "action": "RECORDS"},
            {"key": "settings", "y": 500, "action": "SETTINGS"},
            {"key": "help", "y": 600, "action": "HELP"},
            {"key": "exit", "y": 700, "action": "EXIT"}
        ]
        rects = draw_menu_layout(screen, "title", btns)
        for event in pygame.event.get():
            if event.type == pygame.QUIT: return "EXIT"
            if event.type == pygame.MOUSEBUTTONDOWN:
                act = handle_menu_click(event.pos, rects)
                if act:
                    play_sound("sound_click")
                    return act
        pygame.display.flip()
        clock.tick(60)

def run_settings_state(screen):
    """Меню настроек."""
    clock = pygame.time.Clock()
    while True:
        vol_pct = int(CONFIG["audio"]["volume"] * 100)
        btns = [
            {"key": "vol", "dynamic_val": f"{vol_pct}% (Стрелки)", "y": 300, "action": "VOL"},
            {"key": "lang_btn", "y": 400, "action": "LANG"},
            {"key": "back", "y": 600, "action": "MENU"}
        ]
        rects = draw_menu_layout(screen, "settings", btns)
        
        for event in pygame.event.get():
            if event.type == pygame.QUIT: return "EXIT"
            if event.type == pygame.MOUSEBUTTONDOWN:
                act = handle_menu_click(event.pos, rects)
                if act == "MENU": return "MENU"
                if act == "VOL": adjust_volume(0.1)
                if act == "LANG":
                    CONFIG["settings"]["lang"] = "en" if CONFIG["settings"]["lang"] == "ru" else "ru"
            if event.type == pygame.KEYDOWN:
                if event.key in (pygame.K_RIGHT, pygame.K_UP): adjust_volume(0.1)
                elif event.key in (pygame.K_LEFT, pygame.K_DOWN): adjust_volume(-0.1)
                
        pygame.display.flip()
        clock.tick(60)

def handle_in_game_settings(screen):
    """Мини-меню настроек во время игры (Пауза). Вызывается на ESC."""
    clock = pygame.time.Clock()
    while True:
        vol_pct = int(CONFIG["audio"]["volume"] * 100)
        btns = [
            {"key": "vol", "dynamic_val": f"{vol_pct}%", "y": 300, "action": "VOL"},
            {"key": "back", "y": 500, "action": "BACK"}
        ]
        rects = draw_menu_layout(screen, "settings", btns)
        
        for event in pygame.event.get():
            if event.type == pygame.QUIT: sys.exit()
            if event.type == pygame.KEYDOWN:
                if event.key in (pygame.K_RIGHT, pygame.K_UP): adjust_volume(0.1)
                elif event.key in (pygame.K_LEFT, pygame.K_DOWN): adjust_volume(-0.1)
                elif event.key == pygame.K_ESCAPE: return None # Возврат в игру
            if event.type == pygame.MOUSEBUTTONDOWN:
                act = handle_menu_click(event.pos, rects)
                if act == "VOL": adjust_volume(0.1)
                elif act == "BACK": return None
                
        pygame.display.flip()
        clock.tick(60)

def run_help_state(screen):
    """Экран справки."""
    clock = pygame.time.Clock()
    while True:
        screen.fill(CONFIG["colors"]["bg_menu"])
        draw_text(screen, get_text("help"), 400, 100, 50)
        draw_multiline_text(screen, get_text("rules"), 400, 200, 28)
        draw_text(screen, get_text("back"), 400, 700, 40)
        
        rect = pygame.Rect(250, 675, 300, 50)
        for event in pygame.event.get():
            if event.type == pygame.QUIT: return "EXIT"
            if event.type == pygame.MOUSEBUTTONDOWN and rect.collidepoint(event.pos):
                return "MENU"
        pygame.display.flip()
        clock.tick(60)

def run_records_state(screen):
    """Экран таблицы рекордов."""
    clock, records = pygame.time.Clock(), load_records()
    while True:
        screen.fill(CONFIG["colors"]["bg_menu"])
        draw_text(screen, get_text("records"), 400, 100, 50)
        
        y_pos = 200
        for name, wins in list(records.items())[:10]:
            draw_text(screen, f"{name} ..... {wins}", 400, y_pos, 32)
            y_pos += 45
            
        draw_text(screen, get_text("back") + " (ESC)", 400, 750, 30)
        
        for ev in pygame.event.get():
            if ev.type == pygame.QUIT: return "EXIT"
            if ev.type in (pygame.MOUSEBUTTONDOWN, pygame.KEYDOWN): return "MENU"
            
        pygame.display.flip()
        clock.tick(60)

def run_name_input_state(screen, winner_color, game_surf, angle):
    """
    ЭПИЧНОЕ окно победы.
    Отрисовывается прямо поверх финального положения доски с полупрозрачным затемнением.
    """
    clock = pygame.time.Clock() # Таймер для управления миганием курсора при вводе имени
    name = "" # Строка для хранения вводимого имени игрока
    blink_timer = 0 # Таймер для управления миганием курсора 
    show_cursor = True # Флаг для отображения/скрытия мигающего курсора (нижнее подчеркивание) при вводе имени
    
    # Получаем локализованное название цвета ("Белые" вместо "white")
    win_team_text = get_text(winner_color).upper()
    
    while True:
        # 1. Рисуем замороженную доску на заднем фоне
        draw_rotated_screen(screen, game_surf, angle)
        # 2. Накладываем темное стекло
        draw_overlay(screen, (20, 20, 30), 200)
        
        # 3. Эпичные надписи
        draw_text(screen, win_team_text, 400, 250, 80, CONFIG["colors"]["highlight"])
        draw_text(screen, get_text("win_msg"), 400, 330, 50)
        draw_text(screen, get_text("enter_name"), 400, 480, 35)
        
        # Мигающий курсор для ввода имени
        blink_timer += clock.get_time()
        if blink_timer > 500:
            show_cursor = not show_cursor
            blink_timer = 0
            
        display_name = name + ("_" if show_cursor else "")
        draw_text(screen, display_name, 400, 550, 45, CONFIG["colors"]["highlight"])
        
        draw_text(screen, "ENTER - " + get_text("records"), 400, 750, 25)
        
        # Обработка событий для ввода имени

        for event in pygame.event.get():
            if event.type == pygame.QUIT: return "EXIT"
            if event.type == pygame.KEYDOWN:
                if event.key == pygame.K_RETURN and len(name) > 0:
                    save_winner(name)
                    # При выходе в меню запускаем музыку обратно
                    initialize_audio()
                    return "RECORDS"
                elif event.key == pygame.K_BACKSPACE:
                    name = name[:-1]
                elif event.unicode.isprintable() and len(name) < 15:
                    name += event.unicode
                    
        pygame.display.flip() # Рисуем все изменения на экране
        clock.tick(60) # Ограничиваем до 60 кадров в секунду для плавности и экономии ресурсов

def get_rotated_click(mouse_pos, current_angle):
    """Транслирует клик мыши с учетом угла доски."""
    x, y = mouse_pos # Получаем координаты клика мыши
    if int(current_angle) % 360 == 180:
        x, y = 800 - x, 800 - y # Если доска повернута на 180 градусов, инвертируем координаты клика относительно центра окна
    return y // 100, x // 100 # Преобразуем пиксельные координаты в логические координаты доски (0-7)

def apply_move(board, move_data):
    """Применяет ход логически."""
    start_r, start_c, end_r, end_c = move_data # Распаковываем данные хода
    piece = board.grid[start_r][start_c] # Получаем шашку, которую нужно переместить
    board.move_piece(start_r, start_c, end_r, end_c) # Физически перемещаем шашку на доске (обновляем массив и координаты шашки)
    play_sound("sound_move")
    
    if abs(end_r - start_r) == 2:
        mid_r, mid_c = start_r + (end_r - start_r) // 2, start_c + (end_c - start_c) // 2
        board.remove_piece(mid_r, mid_c)
        
    if piece.color == "white" and end_r == 0: piece.make_king()
    if piece.color == "black" and end_r == 7: piece.make_king()

def run_play_state(screen):
    """Основной игровой цикл."""
    clock, board = pygame.time.Clock(), Board()
    turn, selected, anim = "white", None, {"active": False}
    angle, target_angle = 0.0, 0.0
    
    while True:
        for event in pygame.event.get():
            if event.type == pygame.QUIT: return "EXIT", None
            if event.type == pygame.KEYDOWN and event.key == pygame.K_ESCAPE:
                handle_in_game_settings(screen)
                
            if event.type == pygame.MOUSEBUTTONDOWN and not anim["active"] and angle == target_angle:
                row, col = get_rotated_click(event.pos, angle)
                piece = board.grid[row][col]
                
                if piece and piece.color == turn:
                    selected = (row, col)
                elif selected and is_valid_move(board, selected[0], selected[1], row, col):
                    apply_move(board, (selected[0], selected[1], row, col))
                    anim = {"active": True, "piece": board.grid[row][col], "start_row": selected[0], 
                            "start_col": selected[1], "end_row": row, "end_col": col, "progress": 0.0}
                    selected, turn = None, "black" if turn == "white" else "white"

        if anim["active"]:
            anim["progress"] += CONFIG["animation"]["speed"]
            if anim["progress"] >= 1.0:
                anim["active"], target_angle = False, target_angle + 180
                
        if angle < target_angle:
            angle = min(angle + CONFIG["animation"]["flip_speed"], target_angle)
            
        win = check_win_condition(board)
        
        surf = draw_game_surface(board, selected, anim)
        draw_rotated_screen(screen, surf, angle)
        draw_text(screen, f"ESC - {get_text('settings')}", 680, 25, 20, color=(0, 0, 0))
        
        # ЕСЛИ ПОБЕДА - останавливаем музыку, играем звук победы и передаем финальный кадр доски
        if win and not anim["active"] and angle == target_angle:
            stop_music()
            play_sound("sound_win")
            pygame.display.flip() # Рисуем последний кадр доски перед затемнением
            return "NAME_INPUT", (win, surf, angle) # Передаем кортеж с данными победителя и доски
            
        pygame.display.flip()
        clock.tick(60)

def main():
    """Точка входа."""
    pygame.init()
    screen = pygame.display.set_mode((800, 800))
    pygame.display.set_caption(CONFIG["window"]["title"])
    initialize_audio()
    
    state, win_data = "MENU", None
    while state != "EXIT":
        if state == "MENU": state = run_menu_state(screen)
        elif state == "PLAY": state, win_data = run_play_state(screen)
        elif state == "SETTINGS": state = run_settings_state(screen)
        elif state == "HELP": state = run_help_state(screen)
        elif state == "RECORDS": state = run_records_state(screen)
        elif state == "NAME_INPUT": 
            # win_data теперь содержит (цвет_победителя, картинка_доски, угол_поворота)
            state = run_name_input_state(screen, win_data[0], win_data[1], win_data[2])
            
    pygame.quit()
    sys.exit()

if __name__ == "__main__":
    main()