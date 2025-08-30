/*
 * Copyright (c) 2025 KAnggara75
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * See <https://www.gnu.org/licenses/gpl-3.0.html>.
 *
 * @author KAnggara75 on Sat 30/08/25 13.07
 * @project PakaiWA configs
 * https://github.com/PakaiWA/PakaiWA/tree/main/internal/configs
 */

package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	Pool   *pgxpool.Pool
	App    *fiber.App
	Log    *logrus.Logger
	Config *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	//userRepository := repository.NewUserRepository(config.Log)
	//contactRepository := repository.NewContactRepository(config.Log)
	//addressRepository := repository.NewAddressRepository(config.Log)

	// setup producer
	//var userProducer *messaging.UserProducer
	//var contactProducer *messaging.ContactProducer
	//var addressProducer *messaging.AddressProducer

	//if config.Producer != nil {
	//	userProducer = messaging.NewUserProducer(config.Producer, config.Log)
	//	contactProducer = messaging.NewContactProducer(config.Producer, config.Log)
	//	addressProducer = messaging.NewAddressProducer(config.Producer, config.Log)
	//}

	// setup use cases
	//userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, userProducer)
	//contactUseCase := usecase.NewContactUseCase(config.DB, config.Log, config.Validate, contactRepository, contactProducer)
	//addressUseCase := usecase.NewAddressUseCase(config.DB, config.Log, config.Validate, contactRepository, addressRepository, addressProducer)

	// setup controller
	//userController := http.NewUserController(userUseCase, config.Log)
	//contactController := http.NewContactController(contactUseCase, config.Log)
	//addressController := http.NewAddressController(addressUseCase, config.Log)

	// setup redis & rate limiter
	//redisClient := NewRedis()
	//rateLimiterUtil := util.NewRateLimiterUtil(redisClient)

	// setup middleware
	//authMiddleware := middleware.NewAuth(userUseCase, rateLimiterUtil)

	//routeConfig := routers.RouteConfig{
	//	App:               config.App,
	//	UserController:    userController,
	//	ContactController: contactController,
	//	AddressController: addressController,
	//	AuthMiddleware:    authMiddleware,
	//}
	//routeConfig.Setup()
}
