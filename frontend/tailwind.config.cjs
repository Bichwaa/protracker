/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors')
const defaultTheme = require('tailwindcss/defaultTheme');

export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    "./node_modules/vue-tailwind-wysiwyg-editor/**/*.{js,ts,vue}",
  ],
  theme: {
    colors:{
      ...colors,
      gloom: '#D9D9D9',
      royal: '#490A47',
    },
    extend: {
      fontFamily: {
        sans: ['Poppins', ...defaultTheme.fontFamily.sans], // Use Poppins as the default sans font
      },
    }
  },
  plugins: [],
}