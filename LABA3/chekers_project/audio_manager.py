import pygame
import os
from settings import CONFIG

def initialize_audio():
    """Инициализирует музыку с громкостью из конфига."""
    pygame.mixer.init()
    music_path = CONFIG["audio"]["music"]
    
    if os.path.exists(music_path):
        pygame.mixer.music.load(music_path)
        pygame.mixer.music.set_volume(CONFIG["audio"]["volume"])
        pygame.mixer.music.play(-1)

def play_sound(sound_key):
    """Воспроизводит звук."""
    sound_path = CONFIG["audio"][sound_key]
    if os.path.exists(sound_path):
        sound = pygame.mixer.Sound(sound_path)
        sound.set_volume(CONFIG["audio"]["volume"] + 0.2)
        sound.play()

def adjust_volume(amount):
    """
    Изменяет громкость на заданную величину (например, +0.1 или -0.1).
    Ограничивает значения от 0.0 до 1.0.
    """
    current_vol = CONFIG["audio"]["volume"]
    new_vol = current_vol + amount
    
    # Ограничиваем громкость в пределах 0-100%
    new_vol = max(0.0, min(1.0, new_vol))
    
    CONFIG["audio"]["volume"] = round(new_vol, 2)
    pygame.mixer.music.set_volume(CONFIG["audio"]["volume"])
    
def stop_music():
    """Останавливает фоновую музыку."""
    pygame.mixer.music.stop()