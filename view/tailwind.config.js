/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{html,js}"],
  theme: {
    container:{
      center: true,
      padding: '16px'
    },
    extend: {
      colors: {
        navy: '#3c486b',
        merah: '#f45050',
        kuning: '#f9d949',
        putih: '#f0f0f0',
        second: '#ced3e2'
      },
      fontFamily: {
        firacode: "'Fira Code', monospace;"
      },
      screens: {
        '2xl': '1320px'
      }
    },
  },
  plugins: [],
}

