# Карточка BBB "{{header.title}}"

**[Основные атрибуты](/entities/automated_capability/automated_capability_card?id={{id}})** | [Сценарии и бизнес-функции](/entities/automated_capability/automated_capability_card_page_functions?id={{id}}) |[Автоматизированные системы](/entities/automated_capability/automated_capability_card_page_systems?id={{id}})|[Бизнес-сущности](/entities/automated_capability/automated_capability_card_page_business_entities?id={{id}})| [Бизнес API](/entities/automated_capability/automated_capability_card_page_interfaces?id={{id}}) | [Интеграции](/entities/automated_capability/automated_capability_card_page_integrations?id={{id}}) | [Потребители](/entities/automated_capability/automated_capability_card_page_consumers?id={{id}}) | [Диаграмма окружения](/entities/automated_capability/automated_capability_card_page_environment_diagram?id={{id}}) | [Ошибки ({{errorsCount}})](/entities/automated_capability/automated_capability_card_page_errors?id={{id}})|[Визуализация JSON-схемы BBB](/entities/automated_capability/automated_capability_card_page_bbb_pasport?id={{id}})


## Ключевая информация
| Наименование поля                             | Значение                                     |
|:----------------------------------------------|----------------------------------------------|
| **Результат валидации карточки**              | {{errorsText}}                               | 
| **Идентификатор**                             | {{id}}                                       |  
| **Блок-владелец**                             | {{header.owner_block}}                       |
| **Трайб-владелец**                            | {{header.owner_tribe}}                       |
| **Статус жизненного цикла**                   | {{header.status}}                            |
### Описание  
 {{header.description}}

### Ответственные
| Наименование поля                             | Значение                                                                                                                                                                                                                                  |
|:----------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Владелец**                                  | {{#Employees.Responsibles.Owner}}{{#.}}{{#id_link_owner}}[{{id_owner}}]({{id_link_owner}}){{/id_link_owner}}{{^id_link_owner}}{{id_owner}}{{/id_link_owner}}{{/.}}{{/Employees.Responsibles.Owner}}                                       |
| **Архитектор ДКА**                            | {{#Employees.Responsibles.EEArchitect}}{{#.}}{{#id_link_EEA}}[{{id_EEA}}]({{id_link_EEA}}){{/id_link_EEA}}{{^id_link_EEA}}{{id_EEA}}{{/id_link_EEA}}  {{/.}}{{/Employees.Responsibles.EEArchitect}}                                       |
| **Методолог**                                 | {{#Employees.Responsibles.Methodologist}}{{#.}}{{#id_link_Method}}[{{id_Method}}]({{id_link_Method}}){{/id_link_Method}}{{^id_link_Method}}{{id_Method}}{{/id_link_Method}}  {{/.}}{{/Employees.Responsibles.Methodologist}}              |
| **Solution-Архитектор**                       | {{#Employees.Responsibles.SolArch}}{{#.}}{{#id_link_SolArch}}[{{id_SolArch}}]({{id_link_SolArch}}){{/id_link_SolArch}}{{^id_link_SolArch}}{{id_SolArch}}{{/id_link_SolArch}}  {{/.}}{{/Employees.Responsibles.SolArch}}                   |
| **Команда развития**                          |                                                                                                                                                                                                                                           |
| Ссылка на команду                             | {{#Team.Responsibles.development_team.Team_link}}{{#.}}{{DevTeam_Link}}  {{/.}}{{/Team.Responsibles.development_team.Team_link}}                                                                                                          |
| Участники команды                             | {{#Team.Responsibles.development_team.Team}}{{#.}}{{#id_link_DevTeam}}[{{id_DevTeam}}]({{id_link_DevTeam}}){{/id_link_DevTeam}}{{^id_link_DevTeam}}{{id_DevTeam}}{{/id_link_DevTeam}}  {{/.}}{{/Team.Responsibles.development_team.Team}} |
| **Команда сопровождения**                     |                                                                                                                                                                                                                                           |
| **ИФТ**                                       |                                                                                                                                                                                                                                           |
| Ссылка на команду                             | {{#Team.Responsibles.support_team.IFT.Team_link}}{{#.}}{{IFT_Link}}  {{/.}}{{/Team.Responsibles.support_team.IFT.Team_link}}                                                                                                              |
| Участники команды                             | {{#Team.Responsibles.support_team.IFT.Team}}{{#.}}{{#id_link_IFT}}[{{id_IFT}}]({{id_link_IFT}}){{/id_link_IFT}}{{^id_link_IFT}}{{id_IFT}}{{/id_link_IFT}}  {{/.}}{{/Team.Responsibles.support_team.IFT.Team}}                             |
| **ПСИ**                                       |                                                                                                                                                                                                                                           |
| Ссылка на команду                             | {{#Team.Responsibles.support_team.PSI.Team_link}}{{#.}}{{PSI_Link}}  {{/.}}{{/Team.Responsibles.support_team.PSI.Team_link}}                                                                                                              |
| Участники команды                             | {{#Team.Responsibles.support_team.PSI.Team}}{{#.}}{{#id_link_PSI}}[{{id_PSI}}]({{id_link_PSI}}){{/id_link_PSI}}{{^id_link_PSI}}{{id_PSI}}{{/id_link_PSI}}  {{/.}}{{/Team.Responsibles.support_team.PSI.Team}}                             |
| **ПРОМ**                                      |                                                                                                                                                                                                                                           |
| Ссылка на команду                             | {{#Team.Responsibles.support_team.PROM.Team_link}}{{#.}}{{PROM_Link}}  {{/.}}{{/Team.Responsibles.support_team.PROM.Team_link}}                                                                                                           |
| Участники команды                             | {{#Team.Responsibles.support_team.PROM.Team}}{{#.}}{{#id_link_PROM}}[{{id_PROM}}]({{id_link_PROM}}){{/id_link_PROM}}{{^id_link_PROM}}{{id_PROM}}{{/id_link_PROM}}  {{/.}}{{/Team.Responsibles.support_team.PROM.Team}}                    |


### Архитектурные диаграммы и схемы
| Название | Ссылка  |
|:---------|:----------------------------------|
{{#architecture_schemas}}
{{#.}}
| **{{title}}**  | {{url}} |{{/.}}{{^.}}| ==Не заполнены ссылки на архитектурные диаграммы== |
{{/.}}
{{/architecture_schemas}}
{{^architecture_schemas}}
Отсутствуют ссылки на архитектурные диаграммы
{{/architecture_schemas}}

