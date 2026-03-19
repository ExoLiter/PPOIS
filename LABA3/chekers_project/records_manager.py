import json
import os
from settings import CONFIG

def load_records():
    """Загружает рекорды. Содержит защиту от поврежденного JSON-файла."""
    file_path = CONFIG["paths"]["records_file"]
    if not os.path.exists(file_path):
        return {}
        
    try:
        with open(file_path, "r", encoding="utf-8") as file:
            return json.load(file)
    except Exception:
        # Если файл пустой или сломан, возвращаем пустой словарь, чтобы игра не вылетала
        return {}

def save_winner(player_name):
    """Добавляет победу игроку и сохраняет топ-10."""
    records = load_records()
    records[player_name] = records.get(player_name, 0) + 1
    
    # Сортируем по убыванию побед и берем только первые 10 записей
    sorted_items = sorted(records.items(), key=lambda item: item[1], reverse=True)
    top_ten = dict(sorted_items[:10])
    
    with open(CONFIG["paths"]["records_file"], "w", encoding="utf-8") as file:
        json.dump(top_ten, file, indent=4, ensure_ascii=False) 