from flask import Flask, request, flash, url_for, redirect, render_template, jsonify
from flask_sqlalchemy import SQLAlchemy
import os

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///flaskdb.sqlite3'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
app.config['SECRET_KEY'] = os.uname()[1].split('-')[0]
db = SQLAlchemy(app)

class manPage(db.Model):
    id = db.Column(db.Integer, primary_key = True)    
    topic = db.Column(db.String(50), nullable=False)
    content = db.Column('article_text', db.UnicodeText(), nullable=False)
    source = db.Column('url_source', db.String(50), nullable=True) 

db.create_all()

@app.errorhandler(404)
def pageNotFound(error):
    return render_template("404.html"), 404
app.register_error_handler(404,pageNotFound)

@app.route("/Notes/<page>", methods=['GET'])
def notes(page):
    if page.endswith('.html'):
        page = page[:-5]

    if page.lower() == 'windows':
        return render_template("windows.html")
    elif page.lower() == 'apple':
        return render_template("apple.html")
    elif page.lower() == 'cloud':
        return render_template("cloud.html")
    elif page.lower() == 'osint':
        return render_template("osint.html")
    else:
        return pageNotFound(404)

@app.route('/', methods = ['GET'])
def index():
    return render_template("index.html", title='My-OFFLINE-Forensics-Notes', manPage = manPage.query.all())

@app.route('/new', methods = ['GET','POST'])
def newbie():
    if request.method != 'POST':
        return redirect('404.html')

    newMan = manPage(topic=request.json.get('topic'), content=request.json.get('content'))
    db.session.add(newMan)
    db.session.commit()    
    # print('Another way of accessing: ',request.json['topic'])
    # print('From (User_Agent): ', request.user_agent) # will get useragent
    # print('Data sent: ', request.json.items()) # will print -d of curl cmd 
    # # print('Auth data: ', request.headers.get('Authorization')) # get -H {"Authorization":"x"}
    # print('User data: ', request.authorization.username)# get -u <USERNAME>
    # print('User data: ', request.authorization["password"]) # get -u <username>:<PASSWORD>
    # flash('Man was added successfully!') # use for html version
    # return redirect(url_for('index')) #Will call the index() from the '/' route
    return jsonify('You have successfully added ' + request.json['topic'] + ' to the db!')

if __name__ == '__main__':
    app.run(debug = True)
