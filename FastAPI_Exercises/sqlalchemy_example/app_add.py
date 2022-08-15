# SQLAlchemy: Object Relational Modelling (ORM) for Python, quick-start

# https://www.youtube.com/watch?v=OT5qJBINiJY&ab_channel=JamesSkarzinskas

from sqlalchemy import create_engine, Column, Integer, String, ForeignKey
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker, relationship

Base = declarative_base()

class User(Base):
    __tablename__ = "person"

    id = Column('id', Integer, primary_key=True)
    username = Column('username', String, unique=True)

engine = create_engine('sqlite:///users.db', echo=True)
Base.metadata.create_all(bind=engine)
Session = sessionmaker(bind=engine)

# user = User()
# user.id = 0
# user.username = 'Alice'

# session.add(user)
# session.commit()

session = Session()

users = session.query(User).all()
for user in users:
    print("User with username=%s and id=%d" % (user.username, user.id))

session.close()

