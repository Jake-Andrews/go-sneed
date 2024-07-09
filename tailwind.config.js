/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [ "./**/*.html", "./**/*.templ", "./**/*.go", ],
  theme: {
    extend: {
      colors: {
        bf: {
          50:  "#d9c0d9",
          100: "#b59ec0",
          200: "#6b6591",
          300: "#3b4965",
          400: "#1f3640",
          500: "#0e2224",
          600: "#061212",
          700: "#020808",
          800: "#010505",
          900: "#010407",
        },
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ]
}

