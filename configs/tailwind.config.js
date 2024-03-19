/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './internal/templates/**/*.templ',
    './internal/handler/**/*.go'
  ],
  theme: {
    extend: {
      fontFamily: {
        'inter': ['Inter', 'sans-serif'],
      },
      colors:{
        'light-gray':'#D8DEE9',
        'dark-gray':'#2E3440',
        'light-b': '#ECEFF4'
      }
    }
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography')
  ],
};

