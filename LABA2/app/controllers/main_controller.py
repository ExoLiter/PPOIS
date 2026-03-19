from PyQt6.QtWidgets import QFileDialog, QMessageBox
from app.views.dialogs import StudentDialog, SearchDialog, DeleteDialog

class MainController:
    def __init__(self, manager):
        self.manager = manager
        self.view = None # Главное окно ещё не установлено, будет назначено позже
        self.page_size = 10 # Количество записей на странице для пагинации
        self.current_page = 1 # Текущая страница для отображения данных

    def set_view(self, view):
        # Устанавливаем ссылку на главное окно (view) и обновляем отображение данных
        self.view = view
        self.refresh_view() 

    def add_record(self):
        # Открываем диалог для добавления нового студента и сохраняем данные, если пользователь подтвердил
        dlg = StudentDialog(self.view)
        if dlg.exec(): # Если пользователь нажал "Сохранить" (т.е. диалог был принят) 
            student = dlg.get_data() # Получаем данные студента из диалога
            self.manager.add_student(student) # Добавляем студента в менеджер
            self.refresh_view() # Обновляем отображение данных в главном окне

    def open_search(self): # Открываем диалог для поиска записей по различным критериям
        dlg = SearchDialog(self.view, self.manager) # Передаем ссылку на менеджер, чтобы диалог мог выполнять поиск
        dlg.exec() # Запускаем диалог (он будет работать до тех пор, пока пользователь не закроет его)

    def open_delete(self): # Открываем диалог для удаления записей по различным критериям
        dlg = DeleteDialog(self.view) # Передаем ссылку на менеджер, чтобы диалог мог выполнять удаление
        if dlg.exec(): # Если пользователь подтвердил удаление (т.е. диалог был принят)
            criteria = dlg.get_criteria() # Получаем критерии удаления из диалога
            count = self.manager.delete(criteria) # Удаляем записи, соответствующие критериям, и получаем количество удаленных записей
            QMessageBox.information(self.view, "Удалено", f"Удалено записей: {count}") # Показываем сообщение с количеством удаленных записей
            self.refresh_view() # Обновляем отображение данных в главном окне после удаления

    def gen_data(self): # Генерируем фейковые данные для тестирования (50 записей) и обновляем отображение
        self.manager.generate_fake_data()
        self.refresh_view()

    def save_file(self): 
        fname, _ = QFileDialog.getSaveFileName(self.view, "Сохранить", "", "XML (*.xml)") # Открываем диалог для выбора файла сохранения (только XML)
        if fname: 
            self.manager.save_to_file(fname) # Сохраняем данные в выбранный файл и показываем сообщение об успешном сохранении

    def load_file(self): 
        fname, _ = QFileDialog.getOpenFileName(self.view, "Загрузить", "", "XML (*.xml)") # Открываем диалог для выбора файла загрузки (только XML)
        if fname:
            self.manager.load_from_file(fname) # Загружаем данные из выбранного файла и обновляем отображение данных в главном окне после загрузки
            self.refresh_view() 

    def change_page(self, action):
        all_data = self.manager.get_all()
        total_pages = max(1, (len(all_data) + self.page_size - 1) // self.page_size) # Вычисляем общее количество страниц на основе общего количества записей и размера страницы
        
        if action == 'next' and self.current_page < total_pages: 
            self.current_page += 1 # Если действие - "следующая страница" и текущая страница меньше общего количества страниц, переходим на следующую страницу
        elif action == 'prev' and self.current_page > 1: 
            self.current_page -= 1 # Если действие - "предыдущая страница" и текущая страница больше 1, переходим на предыдущую страницу
        self.refresh_view() # Обновляем отображение данных в главном окне после изменения страницы

    def refresh_view(self): 
        all_data = self.manager.get_all() 
        total_pages = max(1, (len(all_data) + self.page_size - 1) // self.page_size) 
        
        if self.current_page > total_pages: 
            self.current_page = total_pages 
            
        start = (self.current_page - 1) * self.page_size
        page_data = all_data[start : start + self.page_size] 
        
        self.view.update_table(page_data) 
        self.view.update_pagination_label(self.current_page, total_pages)
