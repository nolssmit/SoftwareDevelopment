# Generating Mappings from an Existing MetaData

# https://docs.sqlalchemy.org/en/14/orm/extensions/automap.html

from sqlalchemy import create_engine, MetaData, Table, Column, ForeignKey
from sqlalchemy.ext.automap import automap_base
from sqlalchemy.orm import sessionmaker                           #<---------

engine = create_engine('sqlite:///users.db', echo=True)
Session = sessionmaker(bind=engine)                               #<----------

# produce our own MetaData object
metadata = MetaData()

# we can reflect it ourselves from a database, using options
# such as 'only' to limit what tables we look at...
metadata.reflect(engine, only=['person'])                         #<----------

# we can then produce a set of mappings from this MetaData.
Base = automap_base(metadata=metadata)

# calling prepare() just sets up mapped classes and relationships.
Base.prepare()

# mapped classes are ready
User = Base.classes.person                                        #<-----------

session = Session()

users = session.query(User).all()
for user in users:
    print("User with username=%s and id=%d" % (user.name, user.id))

session.close()