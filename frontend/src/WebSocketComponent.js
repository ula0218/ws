// src/WebSocketComponent.js

import React, { useEffect, useState } from 'react';

const WebSocketComponent = () => {
  const [messages, setMessages] = useState([]); // 狀態：存放所有訊息的陣列
  const [inputMessage, setInputMessage] = useState(''); // 狀態：用來存放輸入框的內容

  // 處理輸入框內容變化的函數
  const handleInputChange = (event) => {
    setInputMessage(event.target.value);
  };

  // 發送訊息的函數
  const sendMessage = () => {
    if (inputMessage.trim() === '') return; // 避免發送空訊息

    const newMessage = {
      message: inputMessage.trim(),
      // 可以根據需要添加其他欄位，例如發送者資訊、時間戳等
    };

    const socket = new WebSocket('ws://localhost:8080/chat');
    socket.onopen = () => {
      console.log('WebSocket 連接已建立。');
      socket.send(JSON.stringify(newMessage)); // 發送訊息
      setMessages((prevMessages) => [...prevMessages, newMessage]); // 更新前端顯示的訊息
      setInputMessage(''); // 清空輸入框
    };

    socket.onclose = () => {
      console.log('WebSocket 連接已關閉。');
    };

    socket.onerror = (error) => {
      console.error('WebSocket 發生錯誤：', error);
    };
  };

  // useEffect 鉤子：處理 WebSocket 連接
  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8080/chat');

    socket.onopen = () => {
      console.log('WebSocket 連接已建立。');
    };

    socket.onmessage = (event) => {
      const message = JSON.parse(event.data);
      console.log('接收到訊息：', message);
      setMessages((prevMessages) => [...prevMessages, message]);
    };

    socket.onclose = () => {
      console.log('WebSocket 連接已關閉。');
    };

    return () => {
      socket.close();
    };
  }, []);

  return (
    <div className="App">
      <h1>WebSocket 聊天室</h1>
      <div className="message-container">
        {messages.map((message, index) => (
          <div key={index} className="message">
            <p>{message.message}</p>
          </div>
        ))}
      </div>
      <div className="input-container">
        <input
          type="text"
          value={inputMessage}
          onChange={handleInputChange}
          placeholder="輸入訊息..."
        />
        <button onClick={sendMessage}>發送</button>
      </div>
    </div>
  );
};

export default WebSocketComponent;
