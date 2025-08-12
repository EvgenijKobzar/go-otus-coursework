workspace {

  model {
    /* Персонажи */
    user = person "Обычный пользователь" "Регистрируется/входит, просматривает каталог, смотрит видео, управляет подпиской."
    admin = person "Администратор" "Управляет контентом, пользователями, тарифами; модерация и аналитика."

    /* Главная система и контейнеры */
    otus = softwareSystem "Otus On-line Films" "Веб и мобильное приложение для просмотра фильмов и сериалов (подписка и разовые покупки)." {

      webUser = container "Web App (User)" "Позволяет пользователям просматривать каталог, смотреть видео, управлять подпиской." "JavaScript/TypeScript, React"

      webAdmin = container "Web App Admin" "Панель управления для администраторов: модерация, управление контентом, аналитика." "JavaScript/TypeScript, React"

      backend = container "Backend API" "Серверная бизнес-логика: аутентификация, каталог, платежи, API для фронтендов." "Java, Spring Boot"

      media = container "Media Processing Service" "Загрузка и обработка видео: транскодинг, генерация превью, публикация в CDN." "Python, FFmpeg"

      db = container "Database" "Хранение данных о пользователях, подписках, каталоге и контенте." "PostgreSQL"

      cache = container "Cache" "Кэширование каталога, данных сессий и токенов." "Redis"
    }

    /* Внешние системы */
    payment = softwareSystem "Платёжный шлюз" "Обрабатывает платежи и подписки (банковские карты, электронные кошельки)." {
        tags "External"
    }

    cdn = softwareSystem "Video CDN" "Хранение и потоковая доставка видео (HLS/DASH)." {
        tags "External"
    }

    idp = softwareSystem "Identity Provider (SSO / OAuth)" "Аутентификация и управление аккаунтами (OAuth2 / SAML)." {
        tags "External"
    }

    mail = softwareSystem "Email Service" "Отправка писем, уведомлений и подтверждений (SMTP / API)." {
        tags "External"
    }

    analytics = softwareSystem "Аналитика / BI" "Сбор и анализ метрик использования и бизнес-отчетов." {
        tags "External"
    }

    /* Связи: Пользователи с контейнерами */
    user -> webUser "Использует" "HTTPS"
    admin -> webAdmin "Использует" "HTTPS"

    /* Связи: Фронтенды с бэкендом */
    webUser -> backend "Отправка запросов (каталог, видео, платежи)" "REST/HTTPS"
    webAdmin -> backend "Отправка запросов (управление, аналитика, модерация)" "REST/HTTPS"

    /* Связи: Бэкенд с внутренними компонентами и внешними системами */
    backend -> db "Чтение и запись данных" "JDBC"
    backend -> cache "Кэширование данных" "Redis protocol"
    backend -> payment "Запросы на оплату / подписку" "REST/HTTPS"
    backend -> cdn "Запросы URL и метаданных видео" "HTTPS"
    backend -> idp "Аутентификация пользователей" "OAuth2 / SAML"
    backend -> mail "Отправка уведомлений" "SMTP / API"
    backend -> analytics "Отправка событий использования" "HTTP events"

    /* Media Processing Service */
    webAdmin -> media "Загрузка исходных видеофайлов" "HTTPS"
    media -> cdn "Публикация транскодированного видео" "HTTPS"
    media -> backend "Отправка метаданных о контенте" "REST/HTTPS"
    media -> db "Сохранение информации о файлах" "JDBC"
  }

  views {
    /* Контекстная диаграмма */
    systemContext otus {
      include *
      autolayout lr
      title "Контекстная диаграмма — Otus On-line Films"
    }

    /* Диаграмма контейнеров */
    container otus {
      include *
      autolayout lr
      title "Container Diagram — Otus On-line Films"
      description "Основные контейнеры системы Otus On-line Films и их взаимодействие между собой и с внешними системами."
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
      element "External" {
        background #999999
        color #000000
      }
    }
  }
}
