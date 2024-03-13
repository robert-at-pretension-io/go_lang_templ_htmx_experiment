/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html"],
  theme: {
    extend: {
      fontFamily: {
        "fraunces": ['Fraunces', 'serif'],
        "playfair": ['Playfair Display', 'serif'],
        "prata": ['Prata', 'serif']
      }
    },
  },
  plugins: [require('@tailwindcss/typography'),],
}

