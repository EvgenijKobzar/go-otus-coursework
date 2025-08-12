workspace {

  model {

    /* Персонажи */
    user = person "Обычный пользователь"
    admin = person "Администратор"

    /* Главная система Otus On-line Films */
    otus = softwareSystem "Otus On-line Films" {
      backend = container "Backend API" "Серверная бизнес-логика" "Java, Spring Boot" {
        /* Компоненты Backend API */
        authService = component "Auth Service" "Проверка JWT"
        catalogService = component "Catalog Service" "Управление каталогом фильмов и сериалов" "Java Service"
        streamingService = component "Streaming Service" "Генерация временных URL и токенов доступа к видео" "Java Service"
        notificationService = component "Notification Service" "Отправка писем и уведомлений" "Java Service"
        analyticsAdapter = component "Analytics Adapter" "Отправка событий в систему аналитики" "Java Service"
        adminService = component "Admin Service" "Администрирование контента, пользователей, модерация" "Java Service"
      }

      db = container "Database" "PostgreSQL"
      cache = container "Cache" "Redis"
      webUser = container "Web App (User)" "Vue3"
      webAdmin = container "Web App Admin" "Vue3"
    }

    /* Внешние системы */
    cdn = softwareSystem "Video CDN" {
        tags "External"
    }
    objectStorage = softwareSystem "Object Storage / CDN (S3)" "Хранение постеров и медиа‑артефактов"{
        tags "External"
    }
    mail = softwareSystem "Email Service" {
        tags "External"
    }
    analytics = softwareSystem "Аналитика / BI" {
        tags "External"
    }

    /* Связи фронтендов с компонентами */
    user -> webUser "Использует" "HTTPS"
    admin -> webAdmin "Использует" "HTTPS"

    webUser -> authService "Вход/регистрация" "HTTPS/REST"
    webUser -> catalogService "Просмотр каталога, поиск" "HTTPS/REST"
    webUser -> streamingService "Запрос на просмотр видео" "HTTPS/REST"

    webAdmin -> authService "Вход" "HTTPS/REST"
    webAdmin -> adminService "Управление контентом, пользователями" "HTTPS/REST"
    webAdmin -> catalogService "Управление каталогом" "HTTPS/REST"

    /* Связи компонентов между собой и с внешними системами */
    adminService -> objectStorage "Работа с файлами"
    streamingService -> cdn "Получение URL и токенов для потокового видео" "HTTPS"
    notificationService -> mail "Отправка email-уведомлений" "SMTP/API"
    analyticsAdapter -> analytics "Передача аналитических событий" "HTTPS"

    /* Доступ к хранилищам */
    catalogService -> db "Чтение/запись данных каталога" "JDBC"
    adminService -> db "Чтение/запись данных пользователей/контента" "JDBC"
    authService -> db "Чтение/запись данных пользователей" "JDBC"

    /* Использование кэша */
    catalogService -> cache "Кэширование данных каталога" "Redis protocol"
    authService -> cache "Кэширование токенов" "Redis protocol"
  }

  views {

    /* Диаграмма контекста */
    systemContext otus {
      include *
      autolayout lr
      title "System Context — Otus On-line Films"
      description "Внешние акторы и системы, взаимодействующие с Otus On-line Films."
    }

    /* Диаграмма контейнеров */
    container otus {
      include *
      autolayout lr
      title "Container Diagram — Otus On-line Films"
      description "Контейнеры системы Otus On-line Films и их взаимодействия."
    }

    /* Диаграмма компонентов Backend API */
    component backend {
      include *
      autolayout lr
      title "Component Diagram — Backend API (Otus On-line Films)"
      description "Ключевые сервисы (компоненты) внутри Backend API и их взаимодействие с фронтендами, хранилищами и внешними системами."
    }

    styles {
      element "Person" {
        shape person
        background #08427b
        color #ffffff
      }
      element "Software System" {
        background #1168bd
        color #ffffff
      }
      element "Container" {
        background #438dd5
        color #ffffff
      }
      element "Component" {
        background #85bbf0
        color #000000
      }
      element "External" {
        background #999999
        color #000000
      }
    }
  }
}
