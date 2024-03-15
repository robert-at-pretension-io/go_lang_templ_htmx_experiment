/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html"],
  theme: {
    extend: {
      fontFamily: {
        "fraunces": ['Fraunces', 'serif'],
        "playfair": ['Playfair Display', 'serif'],
        "prata": ['Prata', 'serif']
      },
      keyframes: {
        wiggle: {
          '0%, 100%': { transform: 'translateX(10%)' },
          '50%': { transform: 'translateX(80%)' },
        },
      },
      animation: {
        'back-and-forth': 'wiggle 4s ease-in-out infinite',
      },
    },
  },
  plugins: [require('@tailwindcss/typography'),],
}

