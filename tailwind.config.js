const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'internal/templates/*.templ',
    'pages/**/*.templ',
  ],
  theme: {
    container: {
      center: true,
      padding: {
        DEFAULT: "1rem",
        mobile: "2rem",
        tablet: "4rem",
        desktop: "5rem",
      },
    },
    extend: {
      colors: {
        primary: colors.indigo,
        secondary: colors.yellow,
        neutral: colors.gray,
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ]
}
