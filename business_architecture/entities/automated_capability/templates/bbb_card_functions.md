| Код бизнес-функции | Заголовок   | Описание бизнес-функции |
|:-------------------|:------------|:------------------------|
{{#business_functions}}
{{#.}}
| {{{bfID}}} | {{bfObject.title}}   |  {{bfObject.description}} |
{{/.}}
{{/business_functions}}
{{^business_functions}}
Бизнес-функции не указаны.
{{/business_functions}}