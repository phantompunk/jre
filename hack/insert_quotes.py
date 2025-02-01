import argparse
import json
from pathlib import Path
import sqlite3


def read_json(file_path):
    with open(file_path, "r") as file:
        return json.load(file)


def write_to_db(json_data, database_file, table_name):
    # Define the schema for the table
    create_table_query = f"""
    CREATE TABLE IF NOT EXISTS {table_name} (
        id TEXT PRIMARY KEY,
        quote TEXT NOT NULL,
        speaker TEXT NOT NULL,
        episode TEXT NOT NULL,
        link TEXT NOT NULL,
        created TEXT NOT NULL,
        start INT NOT NULL,
        score INT NOT NULL
    );
    """
    # Connect to the SQLite database (or create it if it doesn't exist)
    conn = sqlite3.connect(database_file)
    cursor = conn.cursor()

    # Create the table if it doesn't exist
    cursor.execute(create_table_query)

    # Insert data into the table
    for entry in json_data:
        # print(entry.get("id", "NotFound"))
        # if not entry.get("id",""):
        #     print(entry)
        cursor.execute(
            f"""INSERT OR REPLACE INTO {table_name} (id, quote, speaker, episode, link, created, start, score)
            VALUES (?, ?, ?, ?, ?, ?, ?, ?)
            """,
            (
                entry["id"],
                entry["selected"],
                entry["speaker"],
                entry["episode"],
                entry["link"],
                entry["created"],
                entry["start"],
                entry["score"],
            ),
        )

    # Commit the changes and close the connection
    conn.commit()
    conn.close()


def main():
    parser = argparse.ArgumentParser(
        description="Read a JSON file and write its data to a SQLite database table."
    )
    parser.add_argument("file", type=str, help="Path to the JSON file.")
    parser.add_argument(
        "--table",
        type=str,
        default="quotes",
        help="Name of the SQLite table (default: quotes).",
    )
    parser.add_argument(
        "--database",
        type=str,
        default="quotes.db",
        help="Name of the SQLite database file (default: quotes.db).",
    )
    args = parser.parse_args()

    # Validate JSON file
    json_file_path = Path(args.file)
    if not json_file_path.exists():
        print(f"Error: JSON file '{args.file}' not found!")
        return

    # Read the JSON data
    json_data = read_json(json_file_path)

    # Write data to the database
    write_to_db(json_data, args.database, args.table)
    print(f"Data successfully written to {args.database} in the '{args.table}' table.")


if __name__ == "__main__":
    main()
# import hashlib
# from datetime import datetime
#
# template = """INSERT INTO quotes(id, quote, speaker, source, created) VALUES("{id}","{quote}","{speaker}","{source}","{created}");"""
# # template = """INSERT INTO quote({id},{quote});"""
#
# quotes = [
#     "If you don’t think chimps will steal babies and eat them, you haven’t been paying attention to the literature",
#     "In ancient times, the only way you see someone like Brock Lesnar is if he came to your island in a boat, and you ran",
#     "Kindness is one of the best gifts you can bestow…We know that inherently that feels great",
#     "I'm a moron, don't take my advice",
#     "I am the bridge between the meatheads and the potheads",
#     "The quicker we all realize that we've been taught how to live life by people that were operating on the momentum of an ignorant past the quicker we can move to a global ethic of community that doesn't value invented borders or the monopolization of natural resources, but rather the goal of a happier more loving humanity.",
#     "Be the guy you pretend to be when you're trying to get laid.",
#     "Live your life like you're the hero in your own movie",
# ]
#
#
# for quote in quotes:
#     id = hashlib.sha1(quote.encode())
#     print(
#         template.format(
#             id=id.hexdigest()[0:6],
#             quote=quote,
#             speaker="Joe Rogan",
#             source="episode 1",
#             created=datetime.now().isoformat(),
#         )
#     )
