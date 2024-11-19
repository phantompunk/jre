/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.{html,js}"],
  theme: {
    extend: {
      colors: {
        'maroon': '#321C2A',
        'tan': '#F3BB8C',
        'berry': '#E3434B',
      }
    },
  },
  plugins: [],
}

