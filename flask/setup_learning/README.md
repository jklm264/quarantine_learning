# Intro to Setup.py

- [Troubleshooting](#Troubleshooting)

## Easy Setup - Virtualize

`$python3 -m venv [**venv**]` and `$source bin/activate`
- Should create virtualenv within project directory. Ex: `$mkdir -p myproject && python3 -m venv myproject/venv && source myproject/venv/bin/activate`
- This way you can add _venv/_  to `.gitignore` file.

Then run `$python setup.py install` [because jetbrains said so](https://www.jetbrains.com/help/pycharm/creating-and-running-setup-py.html)

## Organizing for distribution

```bash
.
└── setup_learning/  # Python package dir
    ├── packagemeup/  # app source code
    |   ├── __init__.py
    |   ├── run.py
    |   ├── static/
    |   └── templates/
    ├── README.md
    ├── requirements.txt # $pip freeze
    ├── setup.py
    └── venv/
```

[Src](https://www.youtube.com/watch?v=Kg1Yvry_Ydk)

## Troubleshooting

- If pip is not installed: `$python -m ensurepip --default-pip` [src](https://docs.python.org/3/installing/index.html#installing-index)
