# Python Virtualization Techniques

## virtualenv

1. virtualenv -p /bin/python3 ./path_of_virtenv
2. source ./path_of_virtenv/bin/activate
3. deactivate

## pipenv

1. pipenv install --user pipenv
2. pipenv isntall --python 3.6
3. pipenv shell
4. deactivate

## Help

Use the following command if Flask says port already in use: `lsof -i -P -n | grep LISTEN`