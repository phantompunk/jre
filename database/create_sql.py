import hashlib
from datetime import datetime

template = """INSERT INTO quotes(id, quote, speaker, source, created) VALUES("{id}","{quote}","{speaker}","{source}","{created}");"""
# template = """INSERT INTO quote({id},{quote});"""

quotes = [
    "If you don’t think chimps will steal babies and eat them, you haven’t been paying attention to the literature",
    "In ancient times, the only way you see someone like Brock Lesnar is if he came to your island in a boat, and you ran",
    "Kindness is one of the best gifts you can bestow…We know that inherently that feels great",
    "I'm a moron, don't take my advice",
    "I am the bridge between the meatheads and the potheads",
    "The quicker we all realize that we've been taught how to live life by people that were operating on the momentum of an ignorant past the quicker we can move to a global ethic of community that doesn't value invented borders or the monopolization of natural resources, but rather the goal of a happier more loving humanity.",
    "Be the guy you pretend to be when you're trying to get laid.",
    "Live your life like you're the hero in your own movie",
]


for quote in quotes:
    id = hashlib.sha1(quote.encode())
    print(
        template.format(
            id=id.hexdigest()[0:6],
            quote=quote,
            speaker="Joe Rogan",
            source="episode 1",
            created=datetime.now().isoformat(),
        )
    )
