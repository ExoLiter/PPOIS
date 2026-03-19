import pygame
from settings import CONFIG, get_text

def draw_text(surface, text, x_pos, y_pos, size=36, color=None):
    """Универсальная отрисовка одной строки текста."""
    if color is None: color = CONFIG["colors"]["text_main"]
    font = pygame.font.SysFont("arial", size)
    text_surface = font.render(str(text), True, color)
    text_rect = text_surface.get_rect(center=(x_pos, y_pos))
    surface.blit(text_surface, text_rect)

def draw_multiline_text(surface, text, x_pos, start_y, size=24):
    """Отрисовка многострочного текста (для справки)."""
    lines = text.split('\n')
    y_offset = start_y
    for line in lines:
        draw_text(surface, line, x_pos, y_offset, size)
        y_offset += 40

def draw_menu_layout(surface, title_key, buttons_data):
    """Общий шаблон для отрисовки меню и кнопок."""
    surface.fill(CONFIG["colors"]["bg_menu"])
    draw_text(surface, get_text(title_key), 400, 100, 60)
    
    rects = []
    for btn in buttons_data:
        text = get_text(btn["key"])
        if btn.get("dynamic_val") is not None:
            text += str(btn["dynamic_val"])
            
        draw_text(surface, text, 400, btn["y"], 40)
        rect = pygame.Rect(250, btn["y"] - 25, 300, 50)
        rects.append((rect, btn["action"]))
        
    return rects

def handle_menu_click(mouse_pos, buttons_rects):
    """Возвращает действие при клике по кнопке."""
    for rect, action in buttons_rects:
        if rect.collidepoint(mouse_pos):
            return action
    return None

def draw_overlay(surface, color=(0, 0, 0), alpha=180):
    """Рисует полупрозрачный экран поверх текущего кадра."""
    overlay = pygame.Surface((CONFIG["window"]["width"], CONFIG["window"]["height"]))
    overlay.set_alpha(alpha)
    overlay.fill(color)
    surface.blit(overlay, (0, 0))