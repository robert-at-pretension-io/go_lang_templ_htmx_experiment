/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html"],
  theme: {
    extend: {
      fontFamily: {
        "fraunces": ['Fraunces', 'serif']
      }
    },
  },
  plugins: [require('@tailwindcss/typography'),],
}

