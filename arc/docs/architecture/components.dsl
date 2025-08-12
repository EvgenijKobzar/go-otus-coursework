workspace {

  model {

    /* Персонажи */
    user = person "Обычный пользователь"
    admin = person "Администратор"

    /* Главная система Otus On-line Films */
    otus = softwareSystem "Otus On-line Films" {
      backend = container "Backend API" "Серверная бизнес-логика" "Java, Spring Boot" {
        /* Компоненты Backend API */
        authService = component "Auth Service" "Аутентификация и авторизация пользователей" "Spring Security"
        catalogService = component "Catalog Service" "Управление каталогом фильмов и сериалов" "Java Service"
        paymentService = component "Payment Service" "Интеграция с платёжным шлюзом, управление подписками" "Java Service"
        streamingService = component "Streaming Service" "Генерация временных URL и токенов доступа к видео" "Java Service"
        notificationService = component "Notification Service" "Отправка писем и уведомлений" "Java Service"
        analyticsAdapter = component "Analytics Adapter" "Отправка событий в систему аналитики" "Java Service"
        adminService = component "Admin Service" "Администрирование контента, пользователей, модерация" "Java Service"
      }

      db = container "Database" "PostgreSQL"
      cache = container "Cache" "Redis"
      webUser = container "Web App (User)" "React"
      webAdmin = container "Web App Admin" "React"
    }

    /* Внешние системы */
    payment = softwareSystem "Платёжный шлюз" {
      tags "External"
    }
    cdn = softwareSystem "Video CDN" {
      tags "External"
    }
    idp = softwareSystem "Identity Provider (SSO / OAuth)" {
      tags "External"
    }
    mail = softwareSystem "Email Service" {
      tags "External"
    }
    analytics = softwareSystem "Аналитика / BI" {
      tags "External"
    }

    /* Связи фронтендов с компонентами */
    webUser -> authService "Вход/регистрация" "HTTPS/REST"
    webUser -> catalogService "Просмотр каталога, поиск" "HTTPS/REST"
    webUser -> paymentService "Оформление подписки" "HTTPS/REST"
    webUser -> streamingService "Запрос на просмотр видео" "HTTPS/REST"

    webAdmin -> authService "Вход" "HTTPS/REST"
    webAdmin -> adminService "Управление контентом, пользователями" "HTTPS/REST"
    webAdmin -> catalogService "Управление каталогом" "HTTPS/REST"

    /* Связи компонентов между собой и с внешними системами */
    authService -> idp "Аутентификация через SSO" "OAuth2/SAML"
    paymentService -> payment "Запросы на оплату / подписку" "REST/HTTPS"
    streamingService -> cdn "Получение URL и токенов для потокового видео" "HTTPS"
    notificationService -> mail "Отправка email-уведомлений" "SMTP/API"
    analyticsAdapter -> analytics "Передача аналитических событий" "HTTPS"

    /* Доступ к хранилищам */
    catalogService -> db "Чтение/запись данных каталога" "JDBC"
    adminService -> db "Чтение/запись данных пользователей/контента" "JDBC"
    paymentService -> db "Чтение/запись данных подписок" "JDBC"
    authService -> db "Чтение/запись данных пользователей" "JDBC"

    /* Использование кэша */
    catalogService -> cache "Кэширование данных каталога" "Redis protocol"
    authService -> cache "Кэширование токенов" "Redis protocol"
  }

  views {
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
