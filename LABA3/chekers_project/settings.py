import json

def load_configuration():
    """Загружает конфигурационный файл."""
    with open("config.json", "r", encoding="utf-8") as file: 
        return json.load(file) # Переводим JSON в словарь Python

CONFIG = load_configuration() # Загружаем конфигурацию при запуске программы

def get_text(key):
    """
    Возвращает переведенный текст по ключу.
    """
    lang = CONFIG["settings"]["lang"]
    return CONFIG["locales"][lang].get(key, key)