from flask import Flask
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
db = SQLAlchemy(app)


class User(db.Model):
    username = db.column(db.String(64), unique=True)
    password = db.column(db.String(64))
    type = db.column(db.SmallInteger)


class EmptyQuestion(db.Model):
    question_id = db.column(db.Integer, primary_key=True)
    question_name = db.column(db.String(256))
    difficulty = db.column(db.Integer)
    subject = db.column(db.String(64))
    expected_output1 = db.column(db.String(256))
    expected_output2 = db.column(db.String(256))
    expected_output3 = db.column(db.String(256))
    expected_output4 = db.column(db.String(256))
    expected_output5 = db.column(db.String(256))
    expected_output6 = db.column(db.String(256))
    expected_output7 = db.column(db.String(256))
    expected_output8 = db.column(db.String(256))


class EmptyExam(db.Model):
    empty_exam_id = db.column(db.Integer, primary_key=True)
    question1 = db.column(db.Integer, unique=True)
    question2 = db.column(db.Integer, unique=True)
    question3 = db.column(db.Integer, unique=True)
    question4 = db.column(db.Integer, unique=True)
    question5 = db.column(db.Integer, unique=True)
    question6 = db.column(db.Integer, unique=True)
    question7 = db.column(db.Integer, unique=True)
    question8 = db.column(db.Integer, unique=True)
    question9 = db.column(db.Integer, unique=True)
    question10 = db.column(db.Integer, unique=True)
    question11 = db.column(db.Integer, unique=True)
    question12 = db.column(db.Integer, unique=True)
    question13 = db.column(db.Integer, unique=True)
    question14 = db.column(db.Integer, unique=True)
    question15 = db.column(db.Integer, unique=True)
    question16 = db.column(db.Integer, unique=True)
    question17 = db.column(db.Integer, unique=True)
    question18 = db.column(db.Integer, unique=True)
    question19 = db.column(db.Integer, unique=True)
    question20 = db.column(db.Integer, unique=True)


class FinishedQuestion(db.Model):
    student_name = db.column(db.String(64))
    exam_id = db.column(db.Integer)
    question_id = db.column(db.Integer)
    question_name = db.column(db.String(256))
    answer = db.column(db.Text)
    expected_output1 = db.column(db.String(256))
    expected_output2 = db.column(db.String(256))
    expected_output3 = db.column(db.String(256))
    expected_output4 = db.column(db.String(256))
    expected_output5 = db.column(db.String(256))
    expected_output6 = db.column(db.String(256))
    expected_output7 = db.column(db.String(256))
    expected_output8 = db.column(db.String(256))
    temp_score = db.column(db.SmallInteger)
    final_score = db.column(db.SmallInteger)
    comment = db.column(db.String)


class FinishedExam(db.Model):
    student_name = db.column(db.String)
    exam_id = db.column(db.Integer)
    question1 = db.column(db.Integer, unique=True)
    question2 = db.column(db.Integer, unique=True)
    question3 = db.column(db.Integer, unique=True)
    question4 = db.column(db.Integer, unique=True)
    question5 = db.column(db.Integer, unique=True)
    question6 = db.column(db.Integer, unique=True)
    question7 = db.column(db.Integer, unique=True)
    question8 = db.column(db.Integer, unique=True)
    question9 = db.column(db.Integer, unique=True)
    question10 = db.column(db.Integer, unique=True)
    question11 = db.column(db.Integer, unique=True)
    question12 = db.column(db.Integer, unique=True)
    question13 = db.column(db.Integer, unique=True)
    question14 = db.column(db.Integer, unique=True)
    question15 = db.column(db.Integer, unique=True)
    question16 = db.column(db.Integer, unique=True)
    question17 = db.column(db.Integer, unique=True)
    question18 = db.column(db.Integer, unique=True)
    question19 = db.column(db.Integer, unique=True)
    question20 = db.column(db.Integer, unique=True)

