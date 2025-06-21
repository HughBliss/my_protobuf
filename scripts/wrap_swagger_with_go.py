#!/usr/bin/python3
import sys

# Читаем весь ввод из stdin
swagger_content = sys.stdin.read()

# Создаем обертку на Go с корректным экранированием обратных кавычек
go_wrapper = f"""package swagger

import (
    "bytes"
    "text/template"
)

type Meta struct {{
    Title   string
    Version string
    Host    string
}}

func GetSwagger(meta Meta) (string, error) {{
    tmpl, err := template.New("swagger").Parse(templateContent)
    if err != nil {{
        return "", err
    }}

    var buf bytes.Buffer
    if err := tmpl.Execute(&buf, meta); err != nil {{
        return "", err
    }}

    return buf.String(), nil
}}

var templateContent = `{swagger_content.replace('`', '` + "`" + `')}`
"""

# Выводим результат в stdout
sys.stdout.write(go_wrapper)