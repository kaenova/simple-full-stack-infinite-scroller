import React from 'react'

function TextContainer({children}) {
  return (
    <div className="bg-white rounded-md p-3 max-w-4xl shadow-md">
      {children}
    </div>
  )
}

export default TextContainer
