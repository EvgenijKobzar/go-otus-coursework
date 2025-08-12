workspace {

  model {
    /* Персонажи */
    user = person "Клиент" "Регистрируется/входит, просматривает каталог, смотрит видео, создает, обновляет, удаляет сериалы."
    admin = person "Администратор" "Управляет пользователями; модерация."

    /* Главная система */
    otus = softwareSystem "Watch Free Movies Online" "Веб приложение для просмотра и управлением каиталогом сериалов."

    /* Внешние системы */
    // payment = softwareSystem "Платёжный шлюз" "Обрабатывает платежи и подписки (банковские карты, электронные кошельки)."
    cdn = softwareSystem "Video CDN" "Хранение и потоковая доставка видео (HLS/DASH)."{
        tags "External"
    }
    idp = softwareSystem "Identity Provider (SSO / OAuth)" "Аутентификация и управление аккаунтами (OAuth2 / SAML)."{
        tags "External"
    }
    mail = softwareSystem "Email Service" "Отправка писем, уведомлений и подтверждений (SMTP / API)."{
        tags "External"
    }
    analytics = softwareSystem "Аналитика / BI" "Сбор и анализ метрик использования и бизнес-отчетов."{
        tags "External"
    }

    /* Отношения */
    user -> otus "Просмотр каталога, воспроизведение видео, управление каталогом" "HTTPS"
    admin -> otus "Управление контентом, настройками, модерация, просмотр статистики" "HTTPS"

    // otus -> payment "Запросы на оплату / подписку (API)" "REST/HTTPS"
    otus -> cdn "Доставка и трансляция видео (HLS/DASH) — предоставляет URL и токены доступа" "HLS/DASH over HTTPS"
    otus -> idp "Аутентификация пользователей (OAuth2 / SAML)" "OAuth2 / SAML"
    otus -> mail "Отправка уведомлений и подтверждений" "SMTP / Email API"
    otus -> analytics "Отправка событий использования и бизнес-метрик" "HTTP events / batch uploads"
  }

  views {
    systemContext otus {
      include *
      autolayout lr
      title "Контекстная диаграмма — Watch Free Movies Online"
      description "Кто взаимодействует с системой и какие внешние сервисы используются."
    }

    /* По желанию можно добавить view для каждого внешнего элемента отдельно */
    // systemContext payment { include payment, otus; autolayout lr }

    /* Стили */
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
      element "External" {
        background #999999
        color #000000
      }
    }
  }

  /* Метаданные (опционально) */
  // documentation { }
}
