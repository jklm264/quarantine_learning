#!/bin/bash

echo 'Usage: ./script.sh app_name (directory_name'

mcd $1
mkdir templates static

touch app.py config.py requirements.txt #run.py == app.py
touch __init__.py controllers.py models.py #view.py == controllers.py
cd .. # back to root

touch templates/index.html
touch templates/404.html
cd .. #back to root
