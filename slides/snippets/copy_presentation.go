/**
 * @license
 * Copyright Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// [START slides_copy_presentation]
package snippets

import (
	"context"
	"fmt"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/slides/v1"
	"log"
)

func copyPresentation(id string, title string) string {
	ctx := context.Background()
	// Uses env GOOGLE_APPLICATION_CREDENTIALS
	client, err := google.DefaultClient(ctx,
		drive.DriveScope,
		slides.PresentationsScope,
		sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Error creating Google client: %v", err)
	}

	driveService, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Drive client: %v", err)
	}
	// Copy a presentation.
	file := drive.File{
		Title: title,
	}
	presentationCopyFile, err := driveService.Files.Copy(id, &file).Do()
	if err != nil {
		log.Fatalf("Unable to copy presentation. %v", err)
	}
	presentationCopyId := presentationCopyFile.Id
	fmt.Printf("Copied presentation ID: %s", presentationCopyId)
	return presentationCopyId
}

// [END slides_create_presentation]
